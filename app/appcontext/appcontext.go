package appcontext

import (
	"github.com/wafer-bw/udx-discord-bot/app/auth"
	"github.com/wafer-bw/udx-discord-bot/app/config"
	"github.com/wafer-bw/udx-discord-bot/app/handler"
	"github.com/wafer-bw/udx-discord-bot/app/slashcommands"
	"github.com/wafer-bw/udx-discord-bot/app/slashcommands/slashcommand"
)

// AppContext implements `AppContext` properties
type AppContext struct {
	Handler handler.Handler
}

// New returns a new `AppContext` struct
func New(slashCommands ...slashcommand.SlashCommand) *AppContext {
	conf := config.New()
	hndl := handler.New(&handler.Deps{
		Auth:          auth.New(&auth.Deps{}, conf),
		SlashCommands: slashcommands.New(slashCommands),
	}, conf)
	return &AppContext{Handler: hndl}
}
