package chstrat

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/disgoslash/discord"
	"github.com/wafer-bw/udx-discord-bot/common/formulas"
	"github.com/wafer-bw/udx-discord-bot/common/nasdaqapi"
)

var name = "chstrat"
var global = true
var guildIDs = []string{
	"116036580094902275", // UDX
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand instance
var SlashCommand = disgoslash.NewSlashCommand(name, command, chstrat, global, guildIDs)

// command schema for the slash command
var command = &discord.ApplicationCommand{
	Name:        name,
	Description: "Find optimal option calls with an extrinsic risk under 10%",
	Options: []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionTypeString,
			Name:        "Symbol",
			Description: "The symbol/ticker for the underlying. Ex: AMD, TSLA, or YOLO",
			Required:    true,
		},
		{
			Type:        discord.ApplicationCommandOptionTypeString,
			Name:        "Asset Class",
			Description: "The asset class of the underlying",
			Required:    true,
			Choices: []*discord.ApplicationCommandOptionChoice{
				{Name: "Stock", Value: "stocks"},
				{Name: "ETF", Value: "etf"},
			},
		},
	},
}

type viableCallsMap map[string][]*viableCall

type viableCall struct {
	ask              float64
	share            float64
	strike           float64
	extrinsicRisk    float64
	expires          time.Time
	expiresDatestamp string
	expiresReadout   string
	optionURL        string
	greeksURL        string
	content          string
	// delta         float64
}

const targetDelta float64 = 0.75

// const targetDeltaPlusMinus float64 = 0.5

// chstrat - Find optimal option calls with an extrinsic risk under 10%
func chstrat(request *discord.InteractionRequest) (*discord.InteractionResponse, error) {

	symbol := request.Data.Options[0].Value
	assetClass := request.Data.Options[1].Value

	napi := nasdaqapi.NewClient()

	share, err := getSharePrice(napi, symbol, assetClass)
	if err != nil {
		return nil, err
	}

	options, err := napi.GetOptions(symbol, assetClass)
	if err != nil {
		return nil, err
	}

	callsMap, err := getCalls(share, options)
	if err != nil {
		return nil, err
	}

	bestCall, err := getBestCall(napi, callsMap, symbol, assetClass)
	if err != nil {
		return nil, err
	}

	var content string
	if bestCall != nil {
		content = bestCall.content
	} else {
		content = "No valid calls found"
	}

	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: content,
		},
	}, nil
}

func findCallByStrike(calls []*viableCall, strike float64) (*viableCall, bool) {
	for _, call := range calls {
		if call.strike == strike {
			return call, true
		}
	}
	return nil, false
}

func getSharePrice(napi nasdaqapi.ClientInterface, symbol string, assetClass string) (float64, error) {
	quote, err := napi.GetQuote(symbol, assetClass)
	if err != nil {
		return 0, err
	}

	share, err := strconv.ParseFloat(strings.ReplaceAll(quote.Data.PrimaryData.LastSalePrice, "$", ""), 64)
	if err != nil {
		return 0, err
	}
	return share, nil
}

func getCalls(share float64, options *nasdaqapi.OptionsResponse) (viableCallsMap, error) {
	calls := viableCallsMap{}
	expiryGroup := ""
	earliestTargetDate := time.Now().AddDate(0, 0, 99)

	for _, option := range options.Data.Table.Rows {
		var err error
		if option.ExpiryGroup != "" {
			expiryGroup = option.ExpiryGroup
		}
		if expiryGroup == "" {
			continue
		}

		call := &viableCall{share: share}
		call.expires, err = time.Parse("January 02, 2006", expiryGroup)
		if err != nil {
			continue
		}
		if call.expires.Unix() < earliestTargetDate.Unix() {
			continue
		}
		call.expiresDatestamp = call.expires.Format("2006-01-02")
		call.expiresReadout = call.expires.Format("Jan02'06")
		call.optionURL = nasdaqapi.SiteBaseURL + option.URL

		call.strike, err = strconv.ParseFloat(option.Strike, 64)
		if err != nil {
			continue
		}

		call.ask, err = strconv.ParseFloat(option.CallAsk, 64)
		if err != nil {
			continue
		}

		call.extrinsicRisk = formulas.GetExtrinsicRisk(call.share, call.strike, call.ask)
		if call.extrinsicRisk > 10 {
			continue
		}

		calls[call.expiresDatestamp] = append(calls[call.expiresDatestamp], call)
	}
	return calls, nil
}

func getBestCall(napi nasdaqapi.ClientInterface, callsMap viableCallsMap, symbol string, assetClass string) (*viableCall, error) {
	var bestCall *viableCall = nil
	bestMatchValue := float64(1)
	for expiry, calls := range callsMap {
		greeks, err := napi.GetGreeks(symbol, assetClass, expiry)
		if err != nil {
			return nil, err
		}

		for _, greek := range greeks.Data.Table.Rows {
			call, found := findCallByStrike(calls, greek.Strike)
			if !found {
				continue
			}
			if greek.CallDelta > 0.80 || greek.CallDelta < 0.70 { // todo use plusminus
				continue
			}

			score := math.Abs(targetDelta - greek.CallDelta)
			if score < bestMatchValue {
				call.greeksURL = nasdaqapi.SiteBaseURL + greek.URL
				call.content = fmt.Sprintf(
					"%s %.0f CALL Δ%.2f\n%s",
					call.expiresReadout, greek.Strike, greek.CallDelta, call.optionURL)
				bestMatchValue = score
				bestCall = call
			}
		}
	}
	return bestCall, nil
}
