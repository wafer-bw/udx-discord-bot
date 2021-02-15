package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/wafer-bw/udx-discord-bot/commands"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/app"
)

type envVars struct {
	Guilds []string `envconfig:"GUILDS" required:"true" split_words:"true"`
}

var env envVars

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning - %s\n", err)
	}
	if err := envconfig.Process("", &env); err != nil {
		log.Fatalf(err.Error())
	}
}

func getErrsStrings(errs []error) []string {
	errsStr := []string{}
	for _, err := range errs {
		errsStr = append(errsStr, err.Error())
	}
	return errsStr
}

func main() {
	syncer := app.NewSyncer()
	if errs := syncer.Run(env.Guilds, commands.SlashCommandMap); len(errs) > 0 {
		log.Println(strings.Join(getErrsStrings(errs), "\n"))
		os.Exit(1)
	}
	os.Exit(0)
}
