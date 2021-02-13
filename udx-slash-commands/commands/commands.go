package commands

import (
	"github.com/wafer-bw/udx-discord-bot/app/slashcommands/slashcommand"
	"github.com/wafer-bw/udx-discord-bot/udx-slash-commands/commands/extrinsicrisk"
	"github.com/wafer-bw/udx-discord-bot/udx-slash-commands/commands/helloworld"
)

// SlashCommands that the application / bot supports
var SlashCommands = []slashcommand.SlashCommand{
	extrinsicrisk.SlashCommand,
	helloworld.SlashCommand,
}
