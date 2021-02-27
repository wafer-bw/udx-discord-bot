package api

import (
	"net/http"

	"github.com/wafer-bw/disgoslash"
	"github.com/wafer-bw/udx-discord-bot/commands"
	"github.com/wafer-bw/udx-discord-bot/common/creds"
)

// Handler is exported for use as a vercel serverless function
// and acts as the entrypoint for slash command requests
// https://vercel.com/docs/serverless-functions/supported-languages#go
// https://discord.com/developers/docs/interactions/slash-commands#responding-to-an-interaction
func Handler(w http.ResponseWriter, r *http.Request) {
	creds := creds.New()
	handler := &disgoslash.Handler{
		SlashCommandMap: commands.SlashCommandMap,
		Creds:           creds,
	}
	handler.Handle(w, r)
}
