package chstrat

import (
	"fmt"
	"net/http"

	"github.com/wafer-bw/disgoslash/models"
	"github.com/wafer-bw/disgoslash/slashcommands"
	"github.com/wafer-bw/udx-discord-bot/common"
	"github.com/wafer-bw/udx-discord-bot/common/nasdaqapi"
)

var name = "chstrat"
var global = true
var guildIDs = []string{
	"116036580094902275", // UDX
	"810227107967402056", // UDX Bot Dev
}

// SlashCommand instance
var SlashCommand = slashcommands.New(name, command, chstrat, global, guildIDs)

// command schema for the slash command
var command = &models.ApplicationCommand{
	Name:        name,
	Description: "Find optimal option calls with an extrinsic risk under 10%",
	Options: []*models.ApplicationCommandOption{
		{
			Type:        models.ApplicationCommandOptionTypeString,
			Name:        "Symbol",
			Description: "Ticker Symbol ex. AMD, TSLA, YOLO",
			Required:    true,
		},
	},
}

// chstrat - Find optimal option calls with an extrinsic risk under 10%
func chstrat(request *models.InteractionRequest) (*models.InteractionResponse, error) {
	symbol := request.Data.Options[0].Value
	url := nasdaqapi.OptionChainGreeksURL(symbol, "stocks")
	headers := map[string]string{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"}
	status, data, err := common.Request(http.MethodGet, url, headers, nil)
	fmt.Println(status)
	fmt.Println(string(data))
	fmt.Println(err)
	return nil, nil
}
