package tradier

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wafer-bw/udx-discord-bot/common/request"
)

const contentType = "application/json"

// Client for interacting with the tradier API
type Client struct {
	Token    string
	Endpoint string
}

// GetQuote for provided symbol
// https://documentation.tradier.com/brokerage-api/markets/get-quotes
func (client Client) GetQuote(symbol string, greeks bool) (*Quote, error) {
	url := fmt.Sprintf("%s/markets/quotes?symbols=%s&greeks=%t", client.Endpoint, symbol, greeks)
	headers := map[string]string{
		"accept":        contentType,
		"Authorization": fmt.Sprintf("Bearer %s", client.Token),
	}
	status, data, err := request.Do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	} else if status != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", status, data)
	}
	quotes := &QuotesResponse{}
	if err := json.Unmarshal(data, quotes); err != nil {
		return nil, err
	}
	return quotes.Quotes.Quote, nil
}
