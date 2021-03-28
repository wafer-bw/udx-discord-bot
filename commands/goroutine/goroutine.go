package goroutine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/disgoslash/discord"
	"github.com/wafer-bw/udx-discord-bot/common/config"
)

var global = false
var guildIDs = []string{
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand instance
var SlashCommand = disgoslash.NewSlashCommand(command, respond, global, guildIDs)

var command = &discord.ApplicationCommand{
	Name:        "workafter",
	Description: "Test deadline exceeded and work afterwards.",
}

func respond(request *discord.InteractionRequest) *discord.InteractionResponse {
	go continuework(request)

	return &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeAcknowledgeWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: "ACK",
		},
	}
}

func continuework(request *discord.InteractionRequest) {
	time.Sleep(discord.MaxResponseTime * 2)

	conf := config.New()
	token := request.Token
	url := fmt.Sprintf("https://discord.com/api/v8/webhooks/%s/%s", conf.Discord.ClientID, token)

	response := &discord.InteractionResponse{
		Type: discord.InteractionResponseTypeChannelMessageWithSource,
		Data: &discord.InteractionApplicationCommandCallbackData{
			Content: "done",
		},
	}

	body, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(http.Post(url, "application/json", bytes.NewReader(body)))
}
