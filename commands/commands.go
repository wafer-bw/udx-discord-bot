package commands

import (
	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/udx-discord-bot/commands/chstrat"
	"github.com/wafer-bw/udx-discord-bot/commands/extrinsicrisk"
	"github.com/wafer-bw/udx-discord-bot/commands/goroutine"
	"github.com/wafer-bw/udx-discord-bot/commands/helloworld"
	"github.com/wafer-bw/udx-discord-bot/commands/timeout"
)

// SlashCommandMap containing slash commands to be deployed and used live
var SlashCommandMap = disgoslash.NewSlashCommandMap(
	extrinsicrisk.SlashCommand,
	helloworld.SlashCommand,
	chstrat.SlashCommand,
	timeout.SlashCommand,
	goroutine.SlashCommand,
)
