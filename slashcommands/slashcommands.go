package slashcommands

import (
	"github.com/wafer-bw/udx-discord-bot/app/models"
	"github.com/wafer-bw/udx-discord-bot/slashcommands/extrinsicrisk"
)

// SlashCommands that the application / bot supports
var SlashCommands = []models.SlashCommand{
	{Name: extrinsicrisk.Name, Command: extrinsicrisk.Command, Action: extrinsicrisk.Action},
}
