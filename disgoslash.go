package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/docopt/docopt-go"
	"github.com/joho/godotenv"
	"github.com/wafer-bw/udx-discord-bot/app/client"
	"github.com/wafer-bw/udx-discord-bot/app/config"
	"github.com/wafer-bw/udx-discord-bot/app/models"
	"github.com/wafer-bw/udx-discord-bot/app/utils"
)

type appargs struct {
	List        bool   `docopt:"list"`
	Verbose     bool   `docopt:"-v,--verbose"`
	Delete      bool   `docopt:"unregister"`
	Create      bool   `docopt:"register"`
	Global      bool   `docopt:"global"`
	GuildID     string `docopt:"<guildID>"`
	CommandID   string `docopt:"<commandID>"`
	CommandPath string `docopt:"<command-json-path>"`
}

var args *appargs
var cmd client.Client
var usage = `Discord Go Slash Commands

Usage:
  disgoslash list [-v|--verbose]
  disgoslash list <guildID> [-v|--verbose]
  disgoslash unregister <commandID>
  disgoslash unregister <guildID> <commandID>
  disgoslash register <command-json-path>
  disgoslash register <guildID> <command-json-path>
  disgoslash -h | --help

Options:
  -h --help  Show this screen.`

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func parseArgs() *appargs {
	args := &appargs{}
	arguments, err := docopt.ParseDoc(usage)
	if err != nil {
		panic(err)
	}
	if err := arguments.Bind(&args); err != nil {
		panic(err)
	}
	return args
}

func list(cmd client.Client, guildID string, verbose bool) string {
	res, err := cmd.ListApplicationCommands(guildID)
	if err != nil {
		panic(err)
	}
	if verbose {
		return utils.FormatJSON(res)
	}
	output := ""
	for _, command := range res {
		output += fmt.Sprintf("%s - %s: %s\n", command.ID, command.Name, command.Description)
	}
	return output
}

func delete(cmd client.Client, guildID string, commandID string) {
	err := cmd.DeleteApplicationCommand(guildID, commandID)
	if err != nil {
		panic(err)
	}

}

func create(cmd client.Client, guildID string, commandPath string) {
	command := &models.ApplicationCommand{}
	file, err := ioutil.ReadFile(commandPath)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, command); err != nil {
		panic(err)
	}
	if err := cmd.CreateApplicationCommand(guildID, command); err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}

func main() {
	loadEnv()
	cmd = client.New(&client.Deps{}, config.New())
	args = parseArgs()

	if args.List {
		fmt.Println(list(cmd, args.GuildID, args.Verbose))
	} else if args.Delete && args.CommandID != "" {
		delete(cmd, args.GuildID, args.CommandID)
	} else if args.Create && args.CommandPath != "" {
		create(cmd, args.GuildID, args.CommandPath)
	}
}
