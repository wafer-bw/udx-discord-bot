package commands

import (
	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/udx-discord-bot/commands/chstrat"
	"github.com/wafer-bw/udx-discord-bot/commands/extrinsicrisk"
	"github.com/wafer-bw/udx-discord-bot/commands/helloworld"
	"github.com/wafer-bw/udx-discord-bot/commands/verify"
)

// SlashCommandMap containing slash commands to be deployed and used live
var SlashCommandMap = disgoslash.NewSlashCommandMap(
	extrinsicrisk.SlashCommand,
	helloworld.SlashCommand,
	chstrat.SlashCommand,
	verify.SlashCommand,
)
