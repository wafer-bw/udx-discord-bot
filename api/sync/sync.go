package sync

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wafer-bw/udx-discord-bot/api/commands"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/app"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	syncer := app.NewSyncer()
	guilds := []string{
		"116036580094902275", // UDX
		"810227107967402056", // UDX Bot Dev
	}

	if errs := syncer.Run(guilds, commands.SlashCommandMap); len(errs) > 0 {
		os.Exit(1)
	}
	os.Exit(0)
}
