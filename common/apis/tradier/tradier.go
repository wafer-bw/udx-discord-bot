package tradier

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wafer-bw/udx-discord-bot/common/request"
)

const contentType = "application/json"

// Client for interacting with the tradier API
type Client struct {
	Token    string
	Endpoint string
}

// ClientInterface of Client methods
type ClientInterface interface {
	GetQuote(symbol string, greeks bool) (*Quote, error)
}

// GetQuote for provided underlying symbol
// https://documentation.tradier.com/brokerage-api/markets/get-quotes
func (client Client) GetQuote(symbol string, greeks bool) (*Quote, error) {
	url := fmt.Sprintf("%s/markets/quotes?symbols=%s&greeks=%t", client.Endpoint, symbol, greeks)
	headers := map[string]string{"accept": contentType, "Authorization": fmt.Sprintf("Bearer %s", client.Token)}
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

// GetOptionExpirations for provided underlying symbol
// https://documentation.tradier.com/brokerage-api/markets/get-options-expirations
func (client Client) GetOptionExpirations(symbol string, includeAllRoots bool, strikes bool) (*Expirations, error) {
	url := fmt.Sprintf("%s/markets/options/expirations?symbol=%s&includeAllRoots=%t&strikes=%t", client.Endpoint, symbol, includeAllRoots, strikes)
	headers := map[string]string{"accept": contentType, "Authorization": fmt.Sprintf("Bearer %s", client.Token)}
	status, data, err := request.Do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	} else if status != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", status, data)
	}
	expirations := &ExpirationsResponse{}
	if err := json.Unmarshal(data, expirations); err != nil {
		return nil, err
	}
	return expirations.Expirations, nil
}

//GetOptionChain at provided expiration for provided symbol
// https://documentation.tradier.com/brokerage-api/markets/get-options-chains
func (client Client) GetOptionChain(symbol string, expiration string, greeks bool) (Chain, error) {
	url := fmt.Sprintf("%s/markets/options/chains?symbol=%s&expiration=%s&greeks=%t", client.Endpoint, symbol, expiration, greeks)
	headers := map[string]string{"accept": contentType, "Authorization": fmt.Sprintf("Bearer %s", client.Token)}
	status, data, err := request.Do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	} else if status != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", status, data)
	}
	optionChain := &OptionsChainResponse{}
	if err := json.Unmarshal(data, optionChain); err != nil {
		return nil, err
	}
	return optionChain.Options.Chain, nil
}

func getFault(data []byte) error {
	log.Println(string(data))
	fault := &FaultResponse{}
	if err := json.Unmarshal(data, fault); err != nil {
		return err
	}
	return fmt.Errorf("%s (%s)", fault.Fault, fault.Fault.Detail.ErrorCode)
}