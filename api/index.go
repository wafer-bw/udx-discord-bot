package api

import (
	"net/http"

	"github.com/wafer-bw/discobottest/app/appcontext"
	"github.com/wafer-bw/discobottest/app/config"
	"github.com/wafer-bw/discobottest/app/handler"
)

// Handler is exported for use as a vercel serverless function
// and acts as the entrypoint for slash command requests
// https://vercel.com/docs/serverless-functions/supported-languages#go
// https://discord.com/developers/docs/interactions/slash-commands#responding-to-an-interaction
func Handler(w http.ResponseWriter, r *http.Request) {
	conf := config.New()
	app := appcontext.New(&appcontext.Deps{}, conf)
	apiHandler := handler.New(&handler.Deps{Interactions: app.Interactions}, conf)
	apiHandler.Handle(w, r)
}
