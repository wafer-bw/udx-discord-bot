package chstrat

import (
	"bytes"
	"context"
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

type chain struct {
	expiry string
	chain  tradier.Chain
}

type chainResult struct {
	chain chain
	err   error
}

const callsDeadlineDuration = 2850 * time.Millisecond

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
	symbol := *request.Data.Options[0].String

	share, err := getSharePrice(tapi, symbol)
	if err != nil {
		log.Println(err)
		return response(err.Error())
	}

	expirations, err := getExpirations(tapi, symbol, now)
	if err != nil {
		log.Println(err)
		return response(err.Error())
	}

	chains, errs, deadlineExceeded := getChains(tapi, symbol, expirations, now)
	for err := range errs {
		log.Println(err)
	}

	calls, err := getCalls(share, chains)
	if err != nil {
		log.Println(err)
		return response(err.Error())
	}

	bestCalls := getBestCalls(calls)
	sortedBestCalls := sortBestCalls(bestCalls)
	return getResponse(symbol, share, sortedBestCalls, deadlineExceeded)
}

func getSharePrice(tapi tradier.ClientInterface, symbol string) (float64, error) {
	quote, err := tapi.GetQuote(symbol, false)
	if err != nil {
		return 0, err
	}
	return quote.Last, nil
}

func getExpirations(tapi tradier.ClientInterface, symbol string, now time.Time) (tradier.Expirations, error) {
	expirations, err := tapi.GetOptionExpirations(symbol, true, false)
	if err != nil {
		return nil, err
	}

	earliestExpiryDate := now.AddDate(0, 0, 99)
	latestExpiryDate := now.AddDate(0, 0, 365)
	viableExpirations := tradier.Expirations{}
	for _, expiry := range expirations {
		expires, err := time.Parse("2006-01-02", expiry)
		if err != nil {
			log.Println(err)
			continue
		}
		if expires.Unix() < earliestExpiryDate.Unix() || expires.Unix() > latestExpiryDate.Unix() {
			continue
		}
		viableExpirations = append(viableExpirations, expiry)
	}

	return viableExpirations, nil
}

func getChain(chains chan<- chainResult, tapi tradier.ClientInterface, symbol string, expiry string) {
	c, e := tapi.GetOptionChain(symbol, expiry, true)
	res := chainResult{chain: chain{chain: c, expiry: expiry}, err: e}
	chains <- res
}

func getChains(tapi tradier.ClientInterface, symbol string, expirations tradier.Expirations, now time.Time) ([]chain, []error, bool) {
	ctx, cancel := context.WithDeadline(context.Background(), now.Add(callsDeadlineDuration))
	defer cancel()

	chains := make(chan chainResult)
	defer close(chains)

	for _, expiry := range expirations {
		go getChain(chains, tapi, symbol, expiry)
	}

	gatheredErrors := []error{}
	gatheredChains := []chain{}
	for i := 0; i < len(expirations); i++ {
		select {
		case res := <-chains:
			if res.err != nil {
				gatheredErrors = append(gatheredErrors, res.err)
			} else if res.chain.chain != nil {
				gatheredChains = append(gatheredChains, res.chain)
			}
		case <-ctx.Done():
			log.Println(ctx.Err())
			return gatheredChains, gatheredErrors, true
		}
	}
	return gatheredChains, gatheredErrors, false
}

func getCalls(share float64, chains []chain) (viableCallsMap, error) {
	calls := viableCallsMap{}
	for _, chain := range chains {
		expires, err := time.Parse("2006-01-02", chain.expiry)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, option := range chain.chain {
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

			calls[chain.expiry] = append(calls[chain.expiry], &viableCall{
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

func getResponse(symbol string, share float64, bestCalls []*viableCall, deadlineExceeded bool) *discord.InteractionResponse {
	if len(bestCalls) == 0 {
		msg := "No valid calls found"
		if deadlineExceeded {
			msg += "\n_incomplete results due to time limit_"
		}
		return response(msg)
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

	msg := fmt.Sprintf("```\n%s - $%.2f\n%s\n```", symbol, share, buffer.String())
	if deadlineExceeded {
		msg += "\n_incomplete results due to time limit_"
	}

	return response(msg)
}

func response(content string) *discord.InteractionResponse {
	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: content,
		},
	}
}
