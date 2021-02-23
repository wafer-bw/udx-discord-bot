package chstrat

import (
	"encoding/json"
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
	symbol := request.Data.Options[0].Value
	assetClass := request.Data.Options[1].Value

	// get share price
	// get and iterate over options
	// get greeks for options with EV% <= 10
	// return call with closest match of 75% delta

	// options, err := getOptions(symbol, assetClass)
	// if err != nil {
	// 	return nil, err
	// }

	// for _, option := range options.Data.Table.Rows {
	// }

	// for _, filter := range greeks.Data.Filters {
	// 	date, err := time.Parse("2006-01-02", filter.Value)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	// Skip options that are less than 100 days away
	// 	earliestTargetDate := time.Now().AddDate(0, 0, 99)
	// 	if date.Unix() < earliestTargetDate.Unix() {
	// 		continue
	// 	}

	// 	greeks, err := getGreeks(symbol, assetClass, filter.Value)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	for _, optionGreeks := range greeks.Data.Table.Rows {

	// 	}

	// 	// todo - request option chain using `date`
	// 	//        Find calls with delta ~75%, what should +/- be?
	// 	//        Calc call extrinsic risk targeting those <10%
	// 	//        Return matching calls in a message
	// }
	return nil, nil
}

func getGreeks(symbol string, assetClass string, date string) (*nasdaqapi.GreeksResponse, error) {
	url := fmt.Sprintf("https://api.nasdaq.com/api/quote/%s/option-chain/greeks?assetclass=%s", symbol, assetClass)
	if date != "" {
		url += fmt.Sprintf("&date=%s", date)
	}
	headers := map[string]string{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"}
	status, data, err := common.Request(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", status, data)
	}
	greeks := &nasdaqapi.GreeksResponse{}
	if err := json.Unmarshal(data, greeks); err != nil {
		return nil, err
	}
	return greeks, nil
}

func getOptions(symbol string, assetClass string) (*nasdaqapi.OptionsResponse, error) {
	url := fmt.Sprintf("https://api.nasdaq.com/api/quote/%s/option-chain?assetclass=%s&excode=oprac&callput=call&money=at&type=all", symbol, assetClass)
	headers := map[string]string{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"}
	status, data, err := common.Request(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", status, data)
	}
	options := &nasdaqapi.OptionsResponse{}
	if err := json.Unmarshal(data, options); err != nil {
		return nil, err
	}
	return options, nil
}
