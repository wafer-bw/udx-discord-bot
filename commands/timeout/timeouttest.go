package timeout

import (
	"strconv"
	"time"

	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/disgoslash/discord"
)

var global = false
var guildIDs = []string{
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand instance
var SlashCommand = disgoslash.NewSlashCommand(command, timeout, global, guildIDs)

var command = &discord.ApplicationCommand{
	Name:        "testtimeout",
	Description: "Test deadline exceeded.",
	Options: []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionTypeString,
			Name:        "Delay",
			Description: "Delay in ms",
			Required:    true,
		},
	},
}

func timeout(request *discord.InteractionRequest) *discord.InteractionResponse {
	delay, err := strconv.Atoi(request.Data.Options[0].Value)
	if err != nil {
		return &discord.InteractionResponse{
			Type: discord.InteractionResponseTypeChannelMessageWithSource,
			Data: &discord.InteractionApplicationCommandCallbackData{
				Content: err.Error(),
			},
		}
	}

	time.Sleep(time.Duration(delay) * time.Millisecond)

	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: "done",
		},
	}
}
