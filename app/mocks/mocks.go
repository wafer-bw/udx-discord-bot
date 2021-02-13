package mocks

import (
	"github.com/wafer-bw/udx-discord-bot/app/config"
)

// Conf mocks the `config.Config` object
var Conf = &config.Config{
	Credentials: &config.Credentials{
		PublicKey:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		ClientID:     "abc123",
		ClientSecret: "abc123",
		Token:        "abc123",
	},
	DiscordAPI: &config.DiscordAPI{
		BaseURL:    "https://discord.com/api",
		APIVersion: "v8",
	},
}
