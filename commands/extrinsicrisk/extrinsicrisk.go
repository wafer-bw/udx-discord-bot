package extrinsicrisk

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/wafer-bw/udx-discord-bot/disgoslash/models"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/slashcommands"
)

var name = "extrinsicrisk"
var global = false
var guildIDs = []string{
	"116036580094902275", // UDX
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand - the slash command instance
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
	fmt.Println("HERE:")
	log.Println(p.Ask)
	if err != nil {
		fmt.Println(err)
		return &models.InteractionResponse{
			Type: models.InteractionResponseTypeChannelMessageWithSource,
			Data: &models.InteractionApplicationCommandCallbackData{
				Content: "Error parsing command :cry:",
			},
		}, nil
	}

	risk := calcExtrinsicRisk(p)

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

func calcExtrinsicRisk(p *payload) float64 {
	return ((p.Ask - (p.Share - p.Strike)) / p.Share) * 100
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
