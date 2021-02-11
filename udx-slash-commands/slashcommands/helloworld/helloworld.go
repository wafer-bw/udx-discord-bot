package helloworld

import (
	"github.com/wafer-bw/udx-discord-bot/app/commands"
	"github.com/wafer-bw/udx-discord-bot/app/models"
)

var commandName = "hello"

// SlashCommand - the slash command instance
var SlashCommand = commands.NewSlashCommand(commandName, command, Hello)

// command schema for the slash command
var command = &models.ApplicationCommand{
	Name:        commandName,
	Description: "Says hello.",
	Options: []*models.ApplicationCommandOption{
		{
			Type:        models.ApplicationCommandOptionTypeString,
			Name:        "Name",
			Description: "Enter your name",
			Required:    true,
		},
	},
}

// Hello - This is where we put the code to run when a user uses our slash command
func Hello(request *models.InteractionRequest) (*models.InteractionResponse, error) {
	username := request.Data.Options[0].Value
	return &models.InteractionResponse{
		Type: models.InteractionResponseTypeChannelMessageWithSource,
		Data: &models.InteractionApplicationCommandCallbackData{
			Content: "Hello " + username + "!",
		},
	}, nil
}
