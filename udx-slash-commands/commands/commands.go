package commands

import (
	"github.com/wafer-bw/udx-discord-bot/disgoslash/slashcommands"
	"github.com/wafer-bw/udx-discord-bot/udx-slash-commands/commands/extrinsicrisk"
	"github.com/wafer-bw/udx-discord-bot/udx-slash-commands/commands/helloworld"
)

// todo - see if this and the sub directories here can be migrated into the API folder

// SlashCommandMap for the application
var SlashCommandMap = slashcommands.NewMap(
	extrinsicrisk.SlashCommand,
	helloworld.SlashCommand,
)
