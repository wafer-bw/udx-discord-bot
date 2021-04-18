package extrinsicvalue

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/disgoslash/discord"
	"github.com/wafer-bw/udx-discord-bot/common/formulas"
)

var global = false
var guildIDs = []string{
	"116036580094902275", // UDX
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand instance
var SlashCommand = disgoslash.NewSlashCommand(command, extrinsicValue, global, guildIDs)

// command schema for the slash command
var command = &discord.ApplicationCommand{
	Name:        "extrinsicvalue",
	Description: "Calculate an option's extrinsic value percentage using the provided share, strike, & ask prices",
	Options: []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionTypeString,
			Name:        "Share",
			Description: "Share price",
			Required:    true,
		},
		{
			Type:        discord.ApplicationCommandOptionTypeString,
			Name:        "Strike",
			Description: "Strike price",
			Required:    true,
		},
		{
			Type:        discord.ApplicationCommandOptionTypeString,
			Name:        "Ask",
			Description: "Ask price",
			Required:    true,
		},
	},
}

// extrinsicValue - The code which completes the desired action of the slash command.
// Calculate extrinsic risk % for provided `share`, `strike`, & `ask`
func extrinsicValue(request *discord.InteractionRequest) *discord.InteractionResponse {
	p, err := getPayload(request.Data.Options)
	if err != nil {
		log.Println(err)
		return &discord.InteractionResponse{
			Type: discord.InteractionResponseTypeChannelMessageWithSource,
			Data: &discord.InteractionApplicationCommandCallbackData{
				Content: "Error parsing command :cry:",
			},
		}
	}
	risk := formulas.GetExtrinsicValue(p.Share, p.Strike, p.Ask)

	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: fmt.Sprintf("%.2f%%", risk),
		},
	}
}

type payload struct {
	Share  float64
	Strike float64
	Ask    float64
}

func getPayload(options []*discord.ApplicationCommandInteractionDataOption) (*payload, error) {
	if len(options) != 3 {
		return nil, errors.New("missing required options")
	}

	shareStr, ok := options[0].StringValue()
	if !ok {
		return nil, errors.New("missing share price")
	}
	share, err := strconv.ParseFloat(shareStr, 64)
	if err != nil {
		return nil, err
	}

	strikeStr, ok := options[1].StringValue()
	if !ok {
		return nil, errors.New("missing strike price")
	}
	strike, err := strconv.ParseFloat(strikeStr, 64)
	if err != nil {
		return nil, err
	}

	askStr, ok := options[2].StringValue()
	if !ok {
		return nil, errors.New("missing ask price")
	}
	ask, err := strconv.ParseFloat(askStr, 64)
	if err != nil {
		return nil, err
	}
	return &payload{Share: share, Strike: strike, Ask: ask}, nil
}
