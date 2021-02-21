package api

import (
	"net/http"

	"github.com/wafer-bw/disgoslash/handler"
	"github.com/wafer-bw/udx-discord-bot/commands"
)

// Handler is exported for use as a vercel serverless function
// and acts as the entrypoint for slash command requests
// https://vercel.com/docs/serverless-functions/supported-languages#go
// https://discord.com/developers/docs/interactions/slash-commands#responding-to-an-interaction
func Handler(w http.ResponseWriter, r *http.Request) {
	handler := handler.New(commands.SlashCommandMap, nil)
	handler.Handle(w, r)
}
