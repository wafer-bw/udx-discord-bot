package leap

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
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand instance
var SlashCommand = disgoslash.NewSlashCommand(command, leapWrapper, global, guildIDs)

// command schema for the slash command
var command = &discord.ApplicationCommand{
	Name:        "leap",
	Description: "Get the best LEAP per expiry matching provided parameters.",
	Options: []*discord.ApplicationCommandOption{
		{
			Required:    true,
			Name:        "Symbol",
			Description: "The symbol for the underlying. Ex: TSLA",
			Type:        discord.ApplicationCommandOptionTypeString,
		},
		{
			Required:    false,
			Name:        "Target-Delta",
			Description: "Target delta. Default: 75",
			Type:        discord.ApplicationCommandOptionTypeInteger,
		},
		{
			Required:    false,
			Name:        "Min-Delta",
			Description: "Minimum delta. Default: 70",
			Type:        discord.ApplicationCommandOptionTypeInteger,
		},
		{
			Required:    false,
			Name:        "Max-Delta",
			Description: "Maximum delta. Default: 80",
			Type:        discord.ApplicationCommandOptionTypeInteger,
		},
		{
			Required:    false,
			Name:        "Min-DTE",
			Description: "Minimum days to expiry. Default: 99",
			Type:        discord.ApplicationCommandOptionTypeInteger,
		},
		{
			Required:    false,
			Name:        "Max-DTE",
			Description: "Minimum days to expiry. Default: 365",
			Type:        discord.ApplicationCommandOptionTypeInteger,
		},
		{
			Required:    false,
			Name:        "Max-EV-percent",
			Description: "Max extrinsic value percentage. Default: 10",
			Type:        discord.ApplicationCommandOptionTypeInteger,
		},
	},
}

type viableCallsMap map[string][]*viableCall
type bestCallsMap map[string]*viableCall

type viableCall struct {
	strike         float64
	bid            float64
	ask            float64
	extrinsicValue float64
	delta          float64
	expiry         string
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
const defaultMinDelta, defaultTargetDelta, defaultMaxDelta int = 70, 75, 80
const defaultMinDTE, defaultMaxDTE int = 99, 365
const defaultMaxEV int = 10

func leapWrapper(request *discord.InteractionRequest) *discord.InteractionResponse {
	conf := config.New()
	tapi := tradier.New(conf.Tradier)
	return leap(request, tapi, time.Now())
}

func mapIntOptions(options []*discord.ApplicationCommandInteractionDataOption) map[string]int {
	optionsMap := map[string]int{}
	for _, option := range options {
		if val, ok := option.IntValue(); ok {
			optionsMap[option.Name] = val
		}
	}
	return optionsMap
}

// leap - Find optimal option calls with an extrinsic risk under 10%
func leap(request *discord.InteractionRequest, tapi tradier.ClientInterface, now time.Time) *discord.InteractionResponse {
	symbol, _ := request.Data.Options[0].StringValue()

	var ok bool
	var minDelta, targetDelta, maxDelta, minDTE, maxDTE, maxEV int

	optionsMap := mapIntOptions(request.Data.Options)
	log.Println(optionsMap)
	log.Println(request.Data.Options)

	if minDelta, ok = optionsMap["Min-Delta"]; !ok {
		minDelta = defaultMinDelta
	}
	if targetDelta, ok = optionsMap["Target-Delta"]; !ok {
		targetDelta = defaultTargetDelta
	}
	if maxDelta, ok = optionsMap["Max-Delta"]; !ok {
		maxDelta = defaultMaxDelta
	}
	if minDTE, ok = optionsMap["Min-DTE"]; !ok {
		minDTE = defaultMinDTE
	}
	if maxDTE, ok = optionsMap["Max-DTE"]; !ok {
		maxDTE = defaultMaxDTE
	}
	if maxEV, ok = optionsMap["Max-EV"]; !ok {
		maxEV = defaultMaxEV
	}

	if minDelta > targetDelta || maxDelta < targetDelta {
		return response("Please ensure target delta lies between min and max delta")
	}

	share, err := getSharePrice(tapi, symbol)
	if err != nil {
		log.Println(err)
		return response(err.Error())
	}

	expirations, err := getExpirations(tapi, symbol, now, minDTE, maxDTE)
	if err != nil {
		log.Println(err)
		return response(err.Error())
	}

	chains, errs, deadlineExceeded := getChains(tapi, symbol, expirations, now)
	for err := range errs {
		log.Println(err)
	}

	calls, err := getCalls(share, chains, minDelta, maxDelta, maxEV)
	if err != nil {
		log.Println(err)
		return response(err.Error())
	}

	bestCalls := getBestCalls(calls, targetDelta)
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

func getExpirations(tapi tradier.ClientInterface, symbol string, now time.Time, minDTE int, maxDTE int) (tradier.Expirations, error) {
	expirations, err := tapi.GetOptionExpirations(symbol, true, false)
	if err != nil {
		return nil, err
	}

	earliestExpiryDate := now.AddDate(0, 0, minDTE)
	latestExpiryDate := now.AddDate(0, 0, maxDTE)
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

func getCalls(share float64, chains []chain, minDelta int, maxDelta int, maxEV int) (viableCallsMap, error) {
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
			if option.Greeks.Delta > float64(maxDelta)/100 || option.Greeks.Delta < float64(minDelta)/100 {
				continue
			}
			extrinsicValue := formulas.GetExtrinsicValue(share, option.Strike, option.Ask)
			if extrinsicValue > float64(maxEV) {
				continue
			}

			calls[chain.expiry] = append(calls[chain.expiry], &viableCall{
				bid:            option.Bid,
				ask:            option.Ask,
				extrinsicValue: extrinsicValue,
				delta:          option.Greeks.Delta,
				expiry:         expires.Format("Jan02'06"),
				strike:         option.Strike,
			})
		}
	}
	return calls, nil
}

func getBestCalls(callsMap viableCallsMap, targetDelta int) bestCallsMap {
	bestCalls := bestCallsMap{}
	for expiry, calls := range callsMap {
		var bestCall *viableCall = nil
		// how close the call delta is to target delta. Lower is better.
		bestScore := float64(1)
		for _, call := range calls {
			score := math.Abs(float64(targetDelta)/100 - call.delta)
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
			fmt.Sprintf("%.2fER", call.extrinsicValue),
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
