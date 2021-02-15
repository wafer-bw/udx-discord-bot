package commands

import (
	"github.com/wafer-bw/udx-discord-bot/commands/extrinsicrisk"
	"github.com/wafer-bw/udx-discord-bot/commands/helloworld"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/slashcommands"
)

// SlashCommandMap containing slash commands to be deployed and used live
var SlashCommandMap = slashcommands.NewMap(
	extrinsicrisk.SlashCommand,
	helloworld.SlashCommand,
)
