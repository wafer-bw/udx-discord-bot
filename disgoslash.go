package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
	"github.com/joho/godotenv"
	"github.com/wafer-bw/udx-discord-bot/app/client"
	"github.com/wafer-bw/udx-discord-bot/app/config"
	"github.com/wafer-bw/udx-discord-bot/app/errs"
	"github.com/wafer-bw/udx-discord-bot/app/utils"
)

var cmd client.Client
var usage = `Discord Go Slash Commands

Usage:
  disgoslash list [-v|--verbose]
  disgoslash list <guildID> [-v|--verbose]
  disgoslash delete <commandID>
  disgoslash delete <guildID> <commandID>
  disgoslash create <command>
  disgoslash create <guildID> <command>
  disgoslash -h | --help

Options:
  -h --help  Show this screen.`

type args struct {
	List      bool   `docopt:"list"`
	Verbose   bool   `docopt:"-v,--verbose"`
	Delete    bool   `docopt:"delete"`
	Create    bool   `docopt:"create"`
	Global    bool   `docopt:"global"`
	GuildID   string `docopt:"<guildID>"`
	CommandID string `docopt:"<commandID>"`
	Command   string `docopt:"<command>"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	cmd = client.New(&client.Deps{}, config.New())
}

func parseArgs() *args {
	args := &args{}
	arguments, err := docopt.ParseDoc(usage)
	if err != nil {
		panic(err)
	}
	if err := arguments.Bind(&args); err != nil {
		panic(err)
	}
	return args
}

func main() {
	args := parseArgs()

	if args.List {
		res, err := cmd.ListApplicationCommands(args.GuildID)
		if err != nil {
			panic(err)
		}
		if args.Verbose {
			utils.PPrint(res)
		} else {
			for _, command := range res {
				fmt.Printf("%s - %s: %s\n", command.ID, command.Name, command.Description)
			}
		}
	} else if args.Delete && args.CommandID != "" {
		err := cmd.DeleteApplicationCommand(args.GuildID, args.CommandID)
		if err != nil {
			panic(err)
		}
		fmt.Println("Done!")
	} else if args.Create && args.Command != "" {
		panic(errs.ErrNotImplemented)
	}
}
