package appcontext

import (
	"github.com/wafer-bw/udx-discord-bot/app/auth"
	"github.com/wafer-bw/udx-discord-bot/app/commands"
	"github.com/wafer-bw/udx-discord-bot/app/config"
	"github.com/wafer-bw/udx-discord-bot/app/handler"
	"github.com/wafer-bw/udx-discord-bot/app/models"
)

// AppContext implements `AppContext` properties
type AppContext struct {
	Commands commands.Commands
	Handler  handler.Handler
}

// New returns a new `AppContext` struct
func New(slashCommands ...models.SlashCommand) *AppContext {
	conf := config.New()
	auth := auth.New(&auth.Deps{}, conf)
	cmds := commands.New(&commands.Deps{SlashCommands: slashCommands}, conf)
	hndl := handler.New(&handler.Deps{Auth: auth, Commands: cmds}, conf)
	return &AppContext{Commands: cmds, Handler: hndl}
}
