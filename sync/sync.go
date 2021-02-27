package main

import (
	"log"
	"os"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/joho/godotenv"
	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/udx-discord-bot/commands"
	"github.com/wafer-bw/udx-discord-bot/common/creds"
)

type appargs struct {
	GuildIDs []string `docopt:"<guildID>"`
}

var usage = `Sync - Reregister Slash Commands with Discord

Usage:
  sync
  sync <guildID>...
  sync -h | --help

Options:
  -h --help  Show this screen.`

func getErrsStrings(errs []error) []string {
	errsStr := []string{}
	for _, err := range errs {
		errsStr = append(errsStr, err.Error())
	}
	return errsStr
}

func parseArgs() (*appargs, error) {
	args := &appargs{}
	arguments, err := docopt.ParseDoc(usage)
	if err != nil {
		return nil, err
	}
	if err := arguments.Bind(&args); err != nil {
		return nil, err
	}
	return args, nil
}

func init() {
	// Required for running locally but not for workflows
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning - %s\n", err)
	}
}

func main() {
	args, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}
	syncer := &disgoslash.Syncer{
		Creds:           creds.New(),
		SlashCommandMap: commands.SlashCommandMap,
		GuildIDs:        args.GuildIDs,
	}
	if errs := syncer.Sync(); len(errs) > 0 {
		log.Println(strings.Join(getErrsStrings(errs), "\n"))
		os.Exit(1)
	}
}
