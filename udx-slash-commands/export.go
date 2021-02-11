package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/wafer-bw/udx-discord-bot/udx-slash-commands/commands"
)

func main() {
	for _, command := range commands.SlashCommands {
		json, err := json.MarshalIndent(command.Command, "", "    ")
		if err != nil {
			panic(err)
		}
		name := fmt.Sprintf("raw/%s.json", command.Name)
		if err := ioutil.WriteFile(name, json, 0644); err != nil {
			panic(err)
		}
		fmt.Println("Generated", name)
	}
}
