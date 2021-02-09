package mocks

import "github.com/wafer-bw/discobottest/app/config"

// PingRequestBody mocks the body of a ping request
var PingRequestBody = `{"type": 1}`

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
