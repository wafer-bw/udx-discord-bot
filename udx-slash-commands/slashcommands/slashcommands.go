package slashcommands

import (
	"github.com/wafer-bw/udx-discord-bot/app/models"
	"github.com/wafer-bw/udx-discord-bot/udx-slash-commands/slashcommands/extrinsicrisk"
)

// SlashCommands that the application / bot supports
var SlashCommands = []models.SlashCommand{
	extrinsicrisk.SlashCommand,
}
