package commands

import (
	"github.com/wafer-bw/udx-discord-bot/commands/extrinsicrisk"
	"github.com/wafer-bw/udx-discord-bot/commands/helloworld"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/slashcommands"
)

// SlashCommandMap for the application
var SlashCommandMap = slashcommands.NewMap(
	extrinsicrisk.SlashCommand,
	helloworld.SlashCommand,
)
