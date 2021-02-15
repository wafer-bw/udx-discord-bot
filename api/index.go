package api

import (
	"net/http"

	"github.com/wafer-bw/udx-discord-bot/disgoslash/app"
	"github.com/wafer-bw/udx-discord-bot/udx-slash-commands/commands"
)

// Handler is exported for use as a vercel serverless function
// and acts as the entrypoint for slash command requests
// https://vercel.com/docs/serverless-functions/supported-languages#go
// https://discord.com/developers/docs/interactions/slash-commands#responding-to-an-interaction
func Handler(w http.ResponseWriter, r *http.Request) {
	handler := app.NewHandler(commands.SlashCommandMap)
	handler.Handle(w, r)
}
