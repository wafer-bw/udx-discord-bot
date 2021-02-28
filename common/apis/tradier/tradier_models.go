package tradier

type QuotesResponse struct {
	Quotes           *Quotes                `json:"quotes"`
	UnmatchedSymbols map[string]interface{} `json:"unmatched_symbols"`
}

type Quotes struct {
	Quote *Quote `json:"quote"`
}

type Quote struct {
	Symbol           string  `json:"symbol"`            // "AAPL",
	Description      string  `json:"description"`       // "Apple Inc",
	Exchange         string  `json:"exch"`              // "Q",
	Type             string  `json:"type"`              // "stock",
	Last             float64 `json:"last"`              // 121.26,
	Change           float64 `json:"change"`            // 0.27,
	Volume           float64 `json:"volume"`            // 164560390,
	Open             float64 `json:"open"`              // 122.59,
	High             float64 `json:"high"`              // 124.85,
	Low              float64 `json:"low"`               // 121.2,
	Close            float64 `json:"close"`             // 121.26,
	Bid              float64 `json:"bid"`               // 121.69,
	Ask              float64 `json:"ask"`               // 121.76,
	ChangePercentage float64 `json:"change_percentage"` // 0.23,
	AverageVolume    int     `json:"average_volume"`    // 105025607,
	LastVolume       int     `json:"last_volume"`       // 16655878,
	TradeDate        int64   `json:"trade_date"`        // 1614373201223,
	PreviousClose    float64 `json:"prevclose"`         // 120.99,
	Week52High       float64 `json:"week_52_high"`      // 145.09,
	Week52Low        float64 `json:"week_52_low"`       // 53.1525,
	BidSize          int     `json:"bidsize"`           // 2,
	BidExchange      string  `json:"bidexch"`           // "P",
	BidDate          int64   `json:"bid_date"`          // 1614387593000,
	AskSize          int     `json:"asksize"`           // 4,
	AskExchange      string  `json:"askexch"`           // "P",
	AskDate          int64   `json:"ask_date"`          // 1614387599000,
	RootSymbols      string  `json:"root_symbols"`      // "AAPL"
}
