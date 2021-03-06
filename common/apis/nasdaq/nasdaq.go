package nasdaq

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wafer-bw/udx-discord-bot/common/request"
)

// SiteBaseURL is the NASDAQ Website Base URL
const SiteBaseURL = "https://www.nasdaq.com"
const apiBaseURL = "https://api.nasdaq.com"

// Client object
type Client struct {
	apiBaseURL string
}

// ClientInterface methods
type ClientInterface interface {
	GetOptions(symbol string, assetClass string) (*OptionsResponse, error)
	GetQuote(symbol string, assetClass string) (*QuoteResponse, error)
	GetGreeks(symbol string, assetClass string, date string) (*GreeksResponse, error)
}

// NewClient Interface
func NewClient() ClientInterface {
	return construct(apiBaseURL)
}

func construct(apiBaseURL string) ClientInterface {
	return &Client{apiBaseURL: apiBaseURL}
}

// GetOptions chains for provided symbol
func (client *Client) GetOptions(symbol string, assetClass string) (*OptionsResponse, error) {
	url := fmt.Sprintf("%s/api/quote/%s/option-chain?assetclass=%s&excode=oprac&callput=call&money=in&type=all&fromdate=all", client.apiBaseURL, symbol, assetClass)
	headers := map[string]string{
		"X-Forwarded-For": "204.101.161.58",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36",
	}
	response, data, err := request.Do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", response.StatusCode, string(data))
	}
	options := &OptionsResponse{}
	if err := json.Unmarshal(data, options); err != nil {
		return nil, err
	}
	return options, nil
}

// GetQuote for provided symbol
func (client *Client) GetQuote(symbol string, assetClass string) (*QuoteResponse, error) {
	url := fmt.Sprintf("%s/api/quote/%s/info?assetclass=%s", client.apiBaseURL, symbol, assetClass)
	headers := map[string]string{
		"X-Forwarded-For": "204.101.161.58",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36",
	}
	response, data, err := request.Do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", response.StatusCode, data)
	}
	options := &QuoteResponse{}
	if err := json.Unmarshal(data, options); err != nil {
		return nil, err
	}
	return options, nil
}

// GetGreeks for provided symol
func (client *Client) GetGreeks(symbol string, assetClass string, date string) (*GreeksResponse, error) {
	url := fmt.Sprintf("%s/api/quote/%s/option-chain/greeks?assetclass=%s", client.apiBaseURL, symbol, assetClass)
	if date != "" {
		url += fmt.Sprintf("&date=%s", date)
	}
	headers := map[string]string{
		"X-Forwarded-For": "204.101.161.58",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36",
	}
	response, data, err := request.Do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", response.StatusCode, data)
	}
	greeks := &GreeksResponse{}
	if err := json.Unmarshal(data, greeks); err != nil {
		return nil, err
	}
	return greeks, nil
}
