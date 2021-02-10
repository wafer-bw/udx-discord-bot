package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
	"github.com/joho/godotenv"
	"github.com/wafer-bw/udx-discord-bot/app/client"
	"github.com/wafer-bw/udx-discord-bot/app/config"
	"github.com/wafer-bw/udx-discord-bot/app/utils"
)

var cmd client.Client
var usage = `Discord Go Slash Commands

Usage:
  disgoslash list global
  disgoslash list <guilds>...
  disgoslash delete global <commandID>
  disgoslash delete <guildID> <commandID>
  disgoslash -h | --help

Options:
  -h --help                  Show this screen.`

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

	arguments, _ := docopt.ParseDoc(usage)
	fmt.Println(arguments)
}
