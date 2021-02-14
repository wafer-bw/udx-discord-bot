package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/syncer"
	"github.com/wafer-bw/udx-discord-bot/udx-slash-commands/commands"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	syncer := syncer.New()
	guilds := []string{
		"116036580094902275", // UDX
		// "810227107967402056", // UDX Bot Dev
		// "807764305415372810", // Ben Bot Dev
	}

	if err := syncer.Run(guilds, commands.SlashCommandMap); err != nil {
		log.Fatal(err)
	}
}
