package slashcommands

import (
	"strings"

	"github.com/wafer-bw/udx-discord-bot/app/errs"
	"github.com/wafer-bw/udx-discord-bot/app/models"
	"github.com/wafer-bw/udx-discord-bot/app/slashcommands/slashcommand"
)

// SlashCommands map
type SlashCommands map[string]slashcommand.SlashCommand

// New returns a new `SlashCommands` map
func New(slashCommandsSlice []slashcommand.SlashCommand) SlashCommands {
	slashCommands := SlashCommands{} // todo - prepopulate built in slashCommands
	slashCommands.add(slashCommandsSlice...)
	return slashCommands
}

// Run the slash command action for the provided interaction request
func (scs SlashCommands) Run(interaction *models.InteractionRequest) (*models.InteractionResponse, error) {
	command, ok := scs[interaction.Data.Name]
	if !ok {
		return nil, errs.ErrNotImplemented
	}
	return command.Do(interaction)
}

func (scs SlashCommands) add(slashCommandsSlice ...slashcommand.SlashCommand) {
	for _, command := range slashCommandsSlice {
		scs[strings.ToLower(command.Name)] = command
	}
}
