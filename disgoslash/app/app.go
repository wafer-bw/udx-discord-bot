package app

import (
	"github.com/wafer-bw/udx-discord-bot/disgoslash/auth"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/config"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/handler"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/slashcommands"
)

// App implements `App` properties
type App struct {
	Handler handler.Handler
}

// New returns a new `App` struct
func New(slashCommandMap slashcommands.Map) *App {
	conf := config.New()
	hndl := handler.New(&handler.Deps{
		Auth:             auth.New(&auth.Deps{}, conf),
		SlashCommandsMap: slashCommandMap,
	}, conf)
	return &App{Handler: hndl}
}
