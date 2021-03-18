package chstrat

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/disgoslash/discord"
	"github.com/wafer-bw/udx-discord-bot/common/apis/tradier"
	"github.com/wafer-bw/udx-discord-bot/common/config"
	"github.com/wafer-bw/udx-discord-bot/common/formulas"
)

var global = false
var guildIDs = []string{
	"116036580094902275", // UDX
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand instance
var SlashCommand = disgoslash.NewSlashCommand(command, chstratWrapper, global, guildIDs)

// command schema for the slash command
var command = &discord.ApplicationCommand{
	Name:        "chstrat",
	Description: "Get best call per expiry w/ ER<10%, Δ.70-.80 closest to Δ.75, and DTE99<365",
	Options: []*discord.ApplicationCommandOption{
		{
			Required:    true,
			Name:        "symbol",
			Description: "The symbol for the underlying. Ex: TSLA",
			Type:        discord.ApplicationCommandOptionTypeString,
		},
	},
}

type viableCallsMap map[string][]*viableCall
type bestCallsMap map[string]*viableCall

type viableCall struct {
	strike        float64
	bid           float64
	ask           float64
	extrinsicRisk float64
	delta         float64
	expiry        string
}

// todo - move these to inputs from the command with defaults
const targetDelta float64 = 0.75
const minDelta float64 = 0.70
const maxDelta float64 = 0.80

func chstratWrapper(request *discord.InteractionRequest) *discord.InteractionResponse {
	conf := config.New()
	tapi := tradier.New(conf.Tradier)
	return chstrat(request, tapi, time.Now())
}

// chstrat - Find optimal option calls with an extrinsic risk under 10%
func chstrat(request *discord.InteractionRequest, tapi tradier.ClientInterface, now time.Time) *discord.InteractionResponse {
	symbol := request.Data.Options[0].Value

	share, err := getSharePrice(tapi, symbol)
	if err != nil {
		log.Println(err)
		return response(err.Error())
	}

	expirations, err := getExpirations(tapi, symbol)
	if err != nil {
		log.Println(err)
		return response(err.Error())
	}

	calls, err := getCalls(tapi, symbol, share, expirations, now)
	if err != nil {
		log.Println(err)
		return response(err.Error())
	}

	bestCalls := getBestCalls(calls)
	sortedBestCalls := sortBestCalls(bestCalls)
	return getResponse(symbol, share, sortedBestCalls)
}

func getSharePrice(tapi tradier.ClientInterface, symbol string) (float64, error) {
	quote, err := tapi.GetQuote(symbol, false)
	if err != nil {
		return 0, err
	}
	return quote.Last, nil
}

func getExpirations(tapi tradier.ClientInterface, symbol string) (tradier.Expirations, error) {
	expirations, err := tapi.GetOptionExpirations(symbol, true, false)
	if err != nil {
		return nil, err
	}
	return expirations, nil
}

func getCalls(tapi tradier.ClientInterface, symbol string, share float64, expirations tradier.Expirations, now time.Time) (viableCallsMap, error) {
	earliestExpiryDate := now.AddDate(0, 0, 99)
	latestExpiryDate := now.AddDate(0, 0, 365)

	calls := viableCallsMap{}
	for _, expiry := range expirations {
		expires, err := time.Parse("2006-01-02", expiry)
		if err != nil {
			log.Println(err)
			continue
		}
		if expires.Unix() < earliestExpiryDate.Unix() || expires.Unix() > latestExpiryDate.Unix() {
			continue
		}

		chain, err := tapi.GetOptionChain(symbol, expiry, true)
		if err != nil {
			return nil, err
		}
		for _, option := range chain {
			if option.OptionType != tradier.OptionTypeCall {
				continue
			}
			if option.Greeks.Delta > maxDelta || option.Greeks.Delta < minDelta {
				continue
			}
			extrinsicRisk := formulas.GetExtrinsicRisk(share, option.Strike, option.Ask)
			if extrinsicRisk > 10 {
				continue
			}

			calls[expiry] = append(calls[expiry], &viableCall{
				bid:           option.Bid,
				ask:           option.Ask,
				extrinsicRisk: extrinsicRisk,
				delta:         option.Greeks.Delta,
				expiry:        expires.Format("Jan02'06"),
				strike:        option.Strike,
			})
		}
	}
	return calls, nil
}

func getBestCalls(callsMap viableCallsMap) bestCallsMap {
	bestCalls := bestCallsMap{}
	for expiry, calls := range callsMap {
		var bestCall *viableCall = nil
		// how close the call delta is to target delta. Lower is better.
		bestScore := float64(1)
		for _, call := range calls {
			score := math.Abs(targetDelta - call.delta)
			if score < bestScore {
				bestScore = score
				bestCall = call
			}
		}
		bestCalls[expiry] = bestCall
	}
	return bestCalls
}

func sortBestCalls(callsMap bestCallsMap) []*viableCall {
	bestCalls := make([]*viableCall, 0, len(callsMap))
	keys := make([]string, 0, len(callsMap))
	for key := range callsMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		bestCalls = append(bestCalls, callsMap[key])
	}
	return bestCalls
}

func getResponse(symbol string, share float64, bestCalls []*viableCall) *discord.InteractionResponse {
	if len(bestCalls) == 0 {
		return response("No valid calls found")
	}

	rows := []string{}
	buffer := new(bytes.Buffer)
	tabber := tabwriter.NewWriter(buffer, 0, 0, 1, ' ', 0)
	for _, call := range bestCalls {
		row := []string{
			call.expiry,
			fmt.Sprintf("%.2f", call.strike),
			fmt.Sprintf("%.0fΔ", call.delta*100),
			fmt.Sprintf("%.2fER", call.extrinsicRisk),
			fmt.Sprintf("b%.2f", call.bid),
			"-",
			fmt.Sprintf("a%.2f", call.ask),
		}
		rows = append(rows, strings.Join(row, "\t"))
	}
	fmt.Fprint(tabber, strings.Join(rows, "\n"))
	tabber.Flush()

	return response(fmt.Sprintf("```\n%s - $%.2f\n%s\n```", symbol, share, buffer.String()))
}

func response(content string) *discord.InteractionResponse {
	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: content,
		},
	}
}
