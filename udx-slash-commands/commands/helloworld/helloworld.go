package helloworld

import (
	"github.com/wafer-bw/udx-discord-bot/app/models"
	"github.com/wafer-bw/udx-discord-bot/app/slashcommands/slashcommand"
)

var name = "helloworld"
var global = true
var guildIDs = []string{
	"116036580094902275", // UDX
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand - the slash command instance
var SlashCommand = slashcommand.New(name, command, hello, global, guildIDs)

// command schema for the slash command
var command = &models.ApplicationCommand{
	Name:        name,
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

// hello - This is where we put the code to run when a user uses our slash command
func hello(request *models.InteractionRequest) (*models.InteractionResponse, error) {
	username := request.Data.Options[0].Value
	return &models.InteractionResponse{
		Type: models.InteractionResponseTypeChannelMessageWithSource,
		Data: &models.InteractionApplicationCommandCallbackData{
			Content: "Hello " + username + "!",
		},
	}, nil
}
