package commands

import (
	"strings"

	"github.com/wafer-bw/udx-discord-bot/app/config"
	"github.com/wafer-bw/udx-discord-bot/app/errs"
	"github.com/wafer-bw/udx-discord-bot/app/models"
)

// Deps defines `Commands` dependencies
type Deps struct {
	SlashCommands []models.SlashCommand
}

// impl implements `Commands` properties
type impl struct {
	deps          *Deps
	conf          *config.Config
	slashCommands models.SlashCommands
}

// Commands interfaces `Commands` methods
type Commands interface {
	Run(interaction *models.InteractionRequest) (*models.InteractionResponse, error)
}

// New returns a new `Commands` interface
func New(deps *Deps, conf *config.Config) Commands {
	slashCommands := models.SlashCommands{} // todo - prepopulate built in slashCommands
	commandsImpl := &impl{deps: deps, conf: conf, slashCommands: slashCommands}
	commandsImpl.add(deps.SlashCommands...)
	return commandsImpl
}

// Run the slash command action for the provided interaction request
func (impl *impl) Run(interaction *models.InteractionRequest) (*models.InteractionResponse, error) {
	command, ok := impl.slashCommands[interaction.Data.Name]
	if !ok {
		return nil, errs.ErrNotImplemented
	}
	return command.Action(interaction)
}

func (impl *impl) add(slashCommands ...models.SlashCommand) {
	for _, command := range slashCommands {
		key := strings.ToLower(command.Name)
		impl.slashCommands[key] = command
	}
}
