package main

import (
	"github.com/joho/godotenv"
	"github.com/wafer-bw/discobottest/app/config"
	"github.com/wafer-bw/discobottest/app/slashcommands"
	"github.com/wafer-bw/discobottest/app/utils"
)

var com slashcommands.SlashCommands

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	com = slashcommands.New(&slashcommands.Deps{}, config.New())
}

func main() {
	res, err := com.ListGuildApplicationCommands("807764305415372810")
	if err != nil {
		panic(err)
	}
	utils.PPrint(res)
}
