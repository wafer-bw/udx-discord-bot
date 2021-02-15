package app

import (
	"github.com/wafer-bw/udx-discord-bot/disgoslash/auth"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/client"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/config"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/handler"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/slashcommands"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/syncer"
)

var conf *config.Config

// LoadConf loads the configuration object from arg or env
func LoadConf(newConf *config.Config) *config.Config {
	if conf == nil {
		if newConf != nil {
			conf = newConf
		} else {
			conf = config.New()
		}
	}
	return conf
}

// NewHandler returns a new handler interface
func NewHandler(slashCommandMap slashcommands.Map) handler.Handler {
	conf := LoadConf(nil)
	h := handler.New(&handler.Deps{
		Auth:             auth.New(&auth.Deps{}, conf),
		SlashCommandsMap: slashCommandMap,
	}, conf)
	return h
}

// NewSyncer returns a new syncer interface
func NewSyncer() syncer.Syncer {
	conf := LoadConf(nil)
	return syncer.New(&syncer.Deps{
		Client: NewClient(),
	}, conf)
}

// NewClient returns a new client interface
func NewClient() client.Client {
	conf := LoadConf(nil)
	return client.New(&client.Deps{}, conf)
}
