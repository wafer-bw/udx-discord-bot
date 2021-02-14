package appcontext

import (
	"github.com/wafer-bw/udx-discord-bot/disgoslash/auth"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/config"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/handler"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/slashcommands"
)

// AppContext implements `AppContext` properties
type AppContext struct {
	Handler handler.Handler
}

// New returns a new `AppContext` struct
func New(slashCommandMap slashcommands.Map) *AppContext {
	conf := config.New()

	hndl := handler.New(&handler.Deps{
		Auth:             auth.New(&auth.Deps{}, conf),
		SlashCommandsMap: slashCommandMap,
	}, conf)

	return &AppContext{Handler: hndl}
}
