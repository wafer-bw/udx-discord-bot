package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/docopt/docopt-go"
	"github.com/joho/godotenv"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/client"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/config"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/models"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/utils"
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

var cmd client.Client
var usage = `Discord Go Slash Commands

Usage:
  disgoslash list [-v|--verbose]
  disgoslash list <guildID> [-v|--verbose]
  disgoslash unregister <commandID>
  disgoslash unregister <commandID> <guildID>
  disgoslash register <command-json-path>
  disgoslash register <command-json-path> <guildID>
  disgoslash -h | --help

Options:
  -h --help  Show this screen.`

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
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

func list(cmd client.Client, guildID string, verbose bool) (string, error) {
	res, err := cmd.ListApplicationCommands(guildID)
	if err != nil {
		return "", err
	}
	if verbose {
		return utils.FormatJSON(res), nil
	}
	output := ""
	for _, command := range res {
		output += fmt.Sprintf("%s - %s: %s\n", command.ID, command.Name, command.Description)
	}
	return output, nil
}

func delete(cmd client.Client, guildID string, commandID string) error {
	return cmd.DeleteApplicationCommand(guildID, commandID)
}

func create(cmd client.Client, guildID string, command *models.ApplicationCommand) error {
	return cmd.CreateApplicationCommand(guildID, command)
}

func loadCommand(commandPath string) (*models.ApplicationCommand, error) {
	command := &models.ApplicationCommand{}
	file, err := ioutil.ReadFile(commandPath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(file, command); err != nil {
		return nil, err
	}
	return command, nil
}

func main() {
	loadEnv()
	cmd = client.New(&client.Deps{}, config.New())
	args, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	if args.List {
		msg, err := list(cmd, args.GuildID, args.Verbose)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	} else if args.Delete && args.CommandID != "" {
		if err := delete(cmd, args.GuildID, args.CommandID); err != nil {
			log.Fatal(err)
		}
	} else if args.Create && args.CommandPath != "" {
		command, err := loadCommand(args.CommandPath)
		if err != nil {
			log.Fatal(err)
		}
		if err := create(cmd, args.GuildID, command); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Done!")
}
