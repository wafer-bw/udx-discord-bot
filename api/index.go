package api

import (
	"net/http"

	"github.com/wafer-bw/udx-discord-bot/app/appcontext"
	"github.com/wafer-bw/udx-discord-bot/slashcommands"
)

// Handler is exported for use as a vercel serverless function
// and acts as the entrypoint for slash command requests
// https://vercel.com/docs/serverless-functions/supported-languages#go
// https://discord.com/developers/docs/interactions/slash-commands#responding-to-an-interaction
func Handler(w http.ResponseWriter, r *http.Request) {
	app := appcontext.New(slashcommands.SlashCommands...)
	app.Handler.Handle(w, r)
}
