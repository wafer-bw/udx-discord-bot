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
			Type:        discord.ApplicationCommandOptionTypeSubCommandGroup,
			Name:        "subcommandgroupa",
			Description: "GROUP A",
			Options: []*discord.ApplicationCommandOption{
				{
					Name:        "subcommand a",
					Description: "SUB A",
					Type:        discord.ApplicationCommandOptionTypeSubCommand,
					Options: []*discord.ApplicationCommandOption{
						{
							Name:        "string",
							Description: "string",
							Type:        discord.ApplicationCommandOptionTypeString,
							Required:    true,
						},
						{
							Name:        "channel",
							Description: "channel",
							Type:        discord.ApplicationCommandOptionTypeChannel,
							Required:    false,
						},
					},
				},
				{
					Name:        "subcommand b",
					Description: "SUB B",
					Type:        discord.ApplicationCommandOptionTypeSubCommand,
					Options: []*discord.ApplicationCommandOption{
						{
							Name:        "string",
							Description: "string",
							Type:        discord.ApplicationCommandOptionTypeString,
							Required:    true,
						},
					},
				},
			},
		},
		{
			Type:        discord.ApplicationCommandOptionTypeSubCommandGroup,
			Name:        "subcommandgroupb",
			Description: "GROUP B",
			Options: []*discord.ApplicationCommandOption{
				{
					Name:        "subcommand c",
					Description: "SUB C",
					Type:        discord.ApplicationCommandOptionTypeSubCommand,
					Options: []*discord.ApplicationCommandOption{
						{
							Name:        "number",
							Description: "number",
							Type:        discord.ApplicationCommandOptionTypeInteger,
							Required:    true,
						},
					},
				},
				{
					Name:        "subcommand d",
					Description: "SUB D",
					Type:        discord.ApplicationCommandOptionTypeSubCommand,
					Options: []*discord.ApplicationCommandOption{
						{
							Name:        "bool",
							Description: "bool",
							Type:        discord.ApplicationCommandOptionTypeBoolean,
							Required:    true,
						},
					},
				},
			},
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
