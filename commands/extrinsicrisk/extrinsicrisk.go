package extrinsicrisk

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
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand instance
var SlashCommand = disgoslash.NewSlashCommand(command, extrinsicRisk, global, guildIDs)

// command schema for the slash command
var command = &discord.ApplicationCommand{
	Name:        "extrinsicrisk",
	Description: "Calculate an option's extrinsic risk percentage using the provided share, strike, & ask prices",
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

// extrinsicRisk - The code which completes the desired action of the slash command.
// Calculate extrinsic risk % for provided `share`, `strike`, & `ask`
func extrinsicRisk(request *discord.InteractionRequest) *discord.InteractionResponse {
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
	risk := formulas.GetExtrinsicRisk(p.Share, p.Strike, p.Ask)

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

	shareStr, err := options[0].GetString()
	if err != nil {
		return nil, err
	}
	share, err := strconv.ParseFloat(*shareStr, 64)
	if err != nil {
		return nil, err
	}

	strikeStr, err := options[1].GetString()
	if err != nil {
		return nil, err
	}
	strike, err := strconv.ParseFloat(*strikeStr, 64)
	if err != nil {
		return nil, err
	}

	askStr, err := options[1].GetString()
	if err != nil {
		return nil, err
	}
	ask, err := strconv.ParseFloat(*askStr, 64)
	if err != nil {
		return nil, err
	}
	return &payload{Share: share, Strike: strike, Ask: ask}, nil
}
