package commands

import (
	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/udx-discord-bot/commands/chstrat"
	"github.com/wafer-bw/udx-discord-bot/commands/extrinsicvalue"
	"github.com/wafer-bw/udx-discord-bot/commands/leap"
)

// SlashCommandMap containing slash commands to be deployed and used live
var SlashCommandMap = disgoslash.NewSlashCommandMap(
	extrinsicvalue.SlashCommand,
	chstrat.SlashCommand,
	leap.SlashCommand,
)
