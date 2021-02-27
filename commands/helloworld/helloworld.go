package helloworld

import (
	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/disgoslash/discord"
)

var name = "helloworld"
var global = false
var guildIDs = []string{
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand instance
var SlashCommand = disgoslash.NewSlashCommand(name, command, hello, global, guildIDs)

// command schema for the slash command
var command = &discord.ApplicationCommand{
	Name:        name,
	Description: "Says hello.",
	Options: []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionTypeString,
			Name:        "Name",
			Description: "Enter your name",
			Required:    true,
		},
	},
}

// hello - This is where we put the code to run when a user uses our slash command
func hello(request *discord.InteractionRequest) (*discord.InteractionResponse, error) {
	username := request.Data.Options[0].Value
	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: "Hello " + username + "!",
		},
	}, nil
}
