package extrinsicrisk

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/wafer-bw/disgoslash/models"
	"github.com/wafer-bw/disgoslash/slashcommands"
)

var name = "extrinsicrisk"
var global = false
var guildIDs = []string{
	"116036580094902275", // UDX
}

// SlashCommand instance
var SlashCommand = slashcommands.New(name, command, extrinsicRisk, global, guildIDs)

// command schema for the slash command
var command = &models.ApplicationCommand{
	Name:        name,
	Description: "Calculate an option's extrinsic risk percentage using the provided share, strike, & ask prices",
	Options: []*models.ApplicationCommandOption{
		{
			Type:        models.ApplicationCommandOptionTypeString,
			Name:        "Share",
			Description: "Share price",
			Required:    true,
		},
		{
			Type:        models.ApplicationCommandOptionTypeString,
			Name:        "Strike",
			Description: "Strike price",
			Required:    true,
		},
		{
			Type:        models.ApplicationCommandOptionTypeString,
			Name:        "Ask",
			Description: "Ask price",
			Required:    true,
		},
	},
}

// extrinsicRisk - The code which completes the desired action of the slash command.
// Calculate extrinsic risk % for provided `share`, `strike`, & `ask`
func extrinsicRisk(request *models.InteractionRequest) (*models.InteractionResponse, error) {
	p, err := getPayload(request.Data.Options)
	if err != nil {
		log.Println(err)
		return &models.InteractionResponse{
			Type: models.InteractionResponseTypeChannelMessageWithSource,
			Data: &models.InteractionApplicationCommandCallbackData{
				Content: "Error parsing command :cry:",
			},
		}, nil
	}

	risk := getExtrinsicRisk(p.Share, p.Strike, p.Ask)

	return &models.InteractionResponse{
		Type: models.InteractionResponseTypeChannelMessageWithSource,
		Data: &models.InteractionApplicationCommandCallbackData{
			Content: fmt.Sprintf("%.2f%%", risk),
		},
	}, nil
}

type payload struct {
	Share  float64
	Strike float64
	Ask    float64
}

func getPayload(options []*models.ApplicationCommandInteractionDataOption) (*payload, error) {
	if len(options) != 3 {
		return nil, errors.New("missing required options")
	}
	share, err := strconv.ParseFloat(options[0].Value, 64)
	if err != nil {
		return nil, err
	}
	strike, err := strconv.ParseFloat(options[1].Value, 64)
	if err != nil {
		return nil, err
	}
	ask, err := strconv.ParseFloat(options[2].Value, 64)
	if err != nil {
		return nil, err
	}
	return &payload{Share: share, Strike: strike, Ask: ask}, nil
}

func getExtrinsicRisk(share float64, strike float64, ask float64) float64 {
	return ((ask - (share - strike)) / share) * 100
}
