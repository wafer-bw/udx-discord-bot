package verify

import (
	"encoding/json"
	"fmt"

	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/disgoslash/discord"
)

var global = false
var guildIDs = []string{
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand instance
var SlashCommand = disgoslash.NewSlashCommand(command, verify, global, guildIDs)

// command schema for the slash command
var command = &discord.ApplicationCommand{
	Name:        "verify",
	Description: "debugging",
	Options: []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionTypeString,
			Name:        "String",
			Description: "Enter a string",
			Required:    true,
		},
		{
			Type:        discord.ApplicationCommandOptionTypeInteger,
			Name:        "Number",
			Description: "Enter a number",
			Required:    true,
		},
		{
			Type:        discord.ApplicationCommandOptionTypeBoolean,
			Name:        "Boolean",
			Description: "Enter a boolean",
			Required:    true,
		},
		{
			Type:        discord.ApplicationCommandOptionTypeUser,
			Name:        "User",
			Description: "Enter a user",
			Required:    true,
		},
		{
			Type:        discord.ApplicationCommandOptionTypeRole,
			Name:        "Role",
			Description: "Enter a role",
			Required:    true,
		},
	},
}

// verify
func verify(request *discord.InteractionRequest) *discord.InteractionResponse {
	var msg string
	data, err := json.Marshal(request.Data.Options)
	if err != nil {
		msg = err.Error()
	} else {
		msg = string(data)
	}

	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: fmt.Sprintf("```%s```", msg),
		},
	}
}
