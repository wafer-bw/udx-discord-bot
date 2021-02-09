package main

import (
	"github.com/joho/godotenv"
	"github.com/wafer-bw/udx-discord-bot/app/client"
	"github.com/wafer-bw/udx-discord-bot/app/config"
	"github.com/wafer-bw/udx-discord-bot/app/utils"
)

var cmd client.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	cmd = client.New(&client.Deps{}, config.New())
}

func main() {
	res, err := cmd.ListGuildApplicationCommands("807764305415372810")
	if err != nil {
		panic(err)
	}
	utils.PPrint(res)
}
