package tradier

type FaultResponse struct {
	Fault Fault `json:"fault"`
}

type Fault struct {
	Fault  string      `json:"faultstring"`
	Detail FaultDetail `json:"detail"`
}

type FaultDetail struct {
	ErrorCode string `json:"errorcode"`
}

type QuotesResponse struct {
	Quotes           *Quotes                `json:"quotes"`
	UnmatchedSymbols map[string]interface{} `json:"unmatched_symbols"`
}

type Quotes struct {
	Quote *Quote `json:"quote"`
}

type Quote struct {
	Symbol           string    `json:"symbol"`
	Description      string    `json:"description"`
	Exchange         string    `json:"exch"`
	Type             QuoteType `json:"type"`
	Last             float64   `json:"last"`
	Change           float64   `json:"change"`
	Volume           float64   `json:"volume"`
	Open             float64   `json:"open"`
	High             float64   `json:"high"`
	Low              float64   `json:"low"`
	Close            float64   `json:"close"`
	Bid              float64   `json:"bid"`
	Ask              float64   `json:"ask"`
	ChangePercentage float64   `json:"change_percentage"`
	AverageVolume    int       `json:"average_volume"`
	LastVolume       int       `json:"last_volume"`
	TradeDate        int64     `json:"trade_date"`
	PreviousClose    float64   `json:"prevclose"`
	Week52High       float64   `json:"week_52_high"`
	Week52Low        float64   `json:"week_52_low"`
	BidSize          int       `json:"bidsize"`
	BidExchange      string    `json:"bidexch"`
	BidDate          int64     `json:"bid_date"`
	AskSize          int       `json:"asksize"`
	AskExchange      string    `json:"askexch"`
	AskDate          int64     `json:"ask_date"`
	RootSymbols      string    `json:"root_symbols"`
}

type OptionExpirationsResponse struct {
	Expirations *OptionExpirations `json:"expirations"`
}

type OptionExpirations struct {
	Expirations Expirations `json:"date"`
}

type Expirations []string

type Strikes struct {
	Strikes []float64 `json:"strike"`
}

type OptionChainsResponse struct {
	Options *Options `json:"options"`
}

type Options struct {
	Chain Chain `json:"option"`
}

type Chain []*Option

type Option struct {
	Symbol           string     `json:"symbol"`
	Description      string     `json:"description"`
	Exchange         string     `json:"exch"`
	Type             QuoteType  `json:"type"`
	Last             float64    `json:"last"`
	Change           float64    `json:"change"`
	Volume           float64    `json:"volume"`
	Open             float64    `json:"open"`
	High             float64    `json:"high"`
	Low              float64    `json:"low"`
	Close            float64    `json:"close"`
	Bid              float64    `json:"bid"`
	Ask              float64    `json:"ask"`
	UnderlyingSymbol string     `json:"underlying"`
	Strike           float64    `json:"strike"`
	Greeks           *Greeks    `json:"greeks"`
	ChangePercentage float64    `json:"change_percentage"`
	AverageVolume    int        `json:"average_volume"`
	LastVolume       int        `json:"last_volume"`
	TradeDate        int64      `json:"trade_date"`
	PreviousClose    float64    `json:"prevclose"`
	Week52High       float64    `json:"week_52_high"`
	Week52Low        float64    `json:"week_52_low"`
	BidSize          int        `json:"bidsize"`
	BidExchange      string     `json:"bidexch"`
	BidDate          int64      `json:"bid_date"`
	AskSize          int        `json:"asksize"`
	AskExchange      string     `json:"askexch"`
	AskDate          int64      `json:"ask_date"`
	OpenInterest     float64    `json:"open_interest"`
	ContractSize     float64    `json:"contract_size"`
	ExpirationDate   string     `json:"expiration_date"`
	ExpirationType   string     `json:"expiration_type"` // todo enum
	OptionType       OptionType `json:"option_type"`
	RootSymbol       string     `json:"root_symbol"`
}

type Greeks struct {
	Delta     float64 `json:"delta"`
	Gamma     float64 `json:"gamma"`
	Theta     float64 `json:"theta"`
	Vega      float64 `json:"vega"`
	Rho       float64 `json:"rho"`
	Phi       float64 `json:"phi"`
	BidIV     float64 `json:"bid_iv"`
	MidIV     float64 `json:"mid_iv"`
	AskIV     float64 `json:"ask_iv"`
	SMVVol    float64 `json:"smv_vol"`
	UpdatedAt string  `json:"updated_at"`
}

type QuoteType string

const (
	QuoteTypeOption QuoteType = "option"
	QuoteTypeStock  QuoteType = "stock"
	QuoteTypeETF    QuoteType = "etf"
)

type OptionType string

const (
	OptionTypeNil  OptionType = ""
	OptionTypePut  OptionType = "put"
	OptionTypeCall OptionType = "call"
)
