package chstrat

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/disgoslash/discord"
	"github.com/wafer-bw/udx-discord-bot/common/apis/tradier"
	"github.com/wafer-bw/udx-discord-bot/common/config"
	"github.com/wafer-bw/udx-discord-bot/common/formulas"
)

var name = "chstrat"
var global = false
var guildIDs = []string{
	// "116036580094902275", // UDX
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand instance
var SlashCommand = disgoslash.NewSlashCommand(name, command, chstrat, global, guildIDs)

// command schema for the slash command
var command = &discord.ApplicationCommand{
	Name:        name,
	Description: "Find a call w/ an ER under 10% and a Δ between .70-.80 as close to .75 as possible.",
	Options: []*discord.ApplicationCommandOption{
		{
			Required:    true,
			Name:        "symbol",
			Description: "The symbol for the underlying. Ex: TSLA",
			Type:        discord.ApplicationCommandOptionTypeString,
		},
	},
}

type viableCalls []*viableCall

type viableCall struct {
	share         float64
	bid           float64
	ask           float64
	extrinsicRisk float64
	delta         float64
	content       string
}

// todo - move these to inputs from the command with defaults
const targetDelta float64 = 0.75
const minDelta float64 = 0.70
const maxDelta float64 = 0.80

// chstrat - Find optimal option calls with an extrinsic risk under 10%
func chstrat(request *discord.InteractionRequest) (*discord.InteractionResponse, error) {
	symbol := request.Data.Options[0].Value

	conf := config.New()
	tapi := &tradier.Client{Token: conf.Tradier.Token, Endpoint: conf.Tradier.Endpoint}

	share, err := getSharePrice(tapi, symbol)
	if err != nil {
		return nil, err
	}

	expirations, err := getExpirations(tapi, symbol)
	if err != nil {
		return nil, err
	}

	calls, err := getCalls(tapi, symbol, share, expirations)
	if err != nil {
		return nil, err
	}

	bestCall := getBestCall(calls)

	return getResponse(bestCall), nil
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

func getCalls(tapi tradier.ClientInterface, symbol string, share float64, expirations tradier.Expirations) (viableCalls, error) {
	earliestExpiryDate := time.Now().AddDate(0, 0, 99)

	calls := viableCalls{}
	for _, expiry := range expirations {
		expires, err := time.Parse("2006-01-02", expiry)
		if err != nil {
			log.Println(err)
			continue
		}
		if expires.Unix() < earliestExpiryDate.Unix() {
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

			calls = append(calls, &viableCall{
				share:         share,
				bid:           option.Bid,
				ask:           option.Ask,
				extrinsicRisk: extrinsicRisk,
				delta:         option.Greeks.Delta,
				content: fmt.Sprintf(".\n%s\n%.2fΔ %.2fER - Bid: %.2f Ask: %.2f Share: %.2f",
					option.Description, option.Greeks.Delta,
					extrinsicRisk, option.Bid, option.Ask, share,
				),
			})
		}
	}
	return calls, nil
}

func getBestCall(calls viableCalls) *viableCall {
	var bestCall *viableCall = nil
	bestScore := float64(1) // how close the call delta is to target delta.
	for _, call := range calls {
		score := math.Abs(targetDelta - call.delta)
		if score < bestScore {
			bestScore = score
			bestCall = call
		}
	}
	return bestCall
}

func getResponse(bestCall *viableCall) *discord.InteractionResponse {
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
	}
}
