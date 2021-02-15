package commands

import (
	"github.com/wafer-bw/udx-discord-bot/api/commands/extrinsicrisk"
	"github.com/wafer-bw/udx-discord-bot/api/commands/helloworld"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/slashcommands"
)

// todo - see if this and the sub directories here can be migrated into the API folder

// SlashCommandMap for the application
var SlashCommandMap = slashcommands.NewMap(
	extrinsicrisk.SlashCommand,
	helloworld.SlashCommand,
)
