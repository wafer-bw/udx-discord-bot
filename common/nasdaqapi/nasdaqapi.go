package nasdaqapi

import "fmt"

// BaseURL for the NASDAQ API
var BaseURL = "https://api.nasdaq.com/api"

// QuotesURL for the NASDAQ API
func QuotesURL(symbol string) string {
	return fmt.Sprintf("%s/quote/%s", BaseURL, symbol)
}

// OptionChainURL for the NASDAQ API
func OptionChainURL(symbol string) string {
	return fmt.Sprintf("%s/option-chain", QuotesURL(symbol))
}

// OptionChainGreeksURL for the NASDAQ API
func OptionChainGreeksURL(symbol string, assetClass string) string {
	return fmt.Sprintf("%s/greeks?assetclass=%s", OptionChainURL(symbol), assetClass)
}
