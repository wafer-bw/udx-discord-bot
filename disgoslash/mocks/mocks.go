package mocks

import (
	"github.com/wafer-bw/udx-discord-bot/disgoslash/config"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/models"
)

// SlashCommandName mocks a command name
var SlashCommandName = "hello"

// SlashCommandResponseContent mocks a command response message
var SlashCommandResponseContent = "Hello World!"

// InteractionResponse mocks an interaciton response object
var InteractionResponse = &models.InteractionResponse{
	Type: models.InteractionResponseTypeChannelMessageWithSource,
	Data: &models.InteractionApplicationCommandCallbackData{Content: SlashCommandResponseContent},
}

// SlashCommandDo mocks a command `Do` function
func SlashCommandDo(request *models.InteractionRequest) (*models.InteractionResponse, error) {
	return InteractionResponse, nil
}

// Conf mocks the `config.Config` object
var Conf = &config.Config{
	Credentials: &config.Credentials{
		PublicKey: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		ClientID:  "abc123",
		Token:     "abc123",
	},
	DiscordAPI: &config.DiscordAPI{
		BaseURL:    "https://discord.com/api",
		APIVersion: "v8",
	},
}
