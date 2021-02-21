package chstrat

import (
	"encoding/json"
	"errors"
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
			Description: "Ex: AMD, TSLA, or YOLO",
			Required:    true,
		},
		{
			Type:        models.ApplicationCommandOptionTypeString,
			Name:        "Asset Class",
			Description: "",
			Required:    true,
			Choices: []*models.ApplicationCommandOptionChoice{
				{Name: "Stock", Value: "stocks"},
				{Name: "ETF", Value: "etf"},
			},
		},
	},
}

// chstrat - Find optimal option calls with an extrinsic risk under 10%
func chstrat(request *models.InteractionRequest) (*models.InteractionResponse, error) {
	// todo - warning we only have 10 seconds to respond on current vercel free plan
	//        may need to parrellize some work here using goroutines depending on the speed of the nasdaq API
	//        may also need to limit number of chains to request if we start getting rate limited by the nasdaq API by doing the above
	symbol := request.Data.Options[0].Value
	assetClass := request.Data.Options[1].Value
	url := fmt.Sprintf("https://api.nasdaq.com/api/quote/%s/option-chain/greeks?assetclass=%s", symbol, assetClass)
	headers := map[string]string{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"}
	status, data, err := common.Request(http.MethodGet, url, headers, nil)
	if err != nil {
		// todo - handle properly
		return nil, err
	}
	if status != http.StatusOK {
		// todo - handle properly
		return nil, errors.New("!OK STATUS")
	}
	greeks := &nasdaqapi.GreeksResponse{}
	if err := json.Unmarshal(data, greeks); err != nil {
		// todo - handle properly
		return nil, err
	}
	for _, filter := range greeks.Data.Filters {
		fmt.Println(filter.Value)
		// todo - request each option chain that's >100 days out
		//        Find calls with delta ~75%, what should +/- be?
		//        Calc call extrinsic risk targeting those <10%
		//        Return matching calls in a message
	}
	return nil, nil
}
