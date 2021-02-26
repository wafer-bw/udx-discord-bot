package nasdaqapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wafer-bw/udx-discord-bot/common/request"
)

const baseURL = "https://api.nasdaq.com"

// Client object
type Client struct {
	baseURL string
}

// ClientInterface methods
type ClientInterface interface {
	GetOptions(symbol string, assetClass string) (*OptionsResponse, error)
	GetQuote(symbol string, assetClass string) (*QuoteResponse, error)
	GetGreeks(symbol string, assetClass string, date string) (*GreeksResponse, error)
}

// NewClient Interface
func NewClient() ClientInterface {
	return construct(baseURL)
}

func construct(baseURL string) ClientInterface {
	return &Client{baseURL: baseURL}
}

// GetOptions chains for provided symbol
func (client *Client) GetOptions(symbol string, assetClass string) (*OptionsResponse, error) {
	url := fmt.Sprintf("%s/api/quote/%s/option-chain?assetclass=%s&excode=oprac&callput=call&money=in&type=all&fromdate=all", client.baseURL, symbol, assetClass)
	headers := map[string]string{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"}
	status, data, err := request.Do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", status, string(data))
	}
	options := &OptionsResponse{}
	if err := json.Unmarshal(data, options); err != nil {
		return nil, err
	}
	return options, nil
}

// GetQuote for provided symbol
func (client *Client) GetQuote(symbol string, assetClass string) (*QuoteResponse, error) {
	url := fmt.Sprintf("%s/api/quote/%s/info?assetclass=%s", client.baseURL, symbol, assetClass)
	headers := map[string]string{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"}
	status, data, err := request.Do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", status, data)
	}
	options := &QuoteResponse{}
	if err := json.Unmarshal(data, options); err != nil {
		return nil, err
	}
	return options, nil
}

// GetGreeks for provided symol
func (client *Client) GetGreeks(symbol string, assetClass string, date string) (*GreeksResponse, error) {
	url := fmt.Sprintf("%s/api/quote/%s/option-chain/greeks?assetclass=%s", client.baseURL, symbol, assetClass)
	if date != "" {
		url += fmt.Sprintf("&date=%s", date)
	}
	headers := map[string]string{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"}
	status, data, err := request.Do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", status, data)
	}
	greeks := &GreeksResponse{}
	if err := json.Unmarshal(data, greeks); err != nil {
		return nil, err
	}
	return greeks, nil
}
