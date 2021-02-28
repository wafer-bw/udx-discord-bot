package nasdaq

// ResponseStatus of request
type ResponseStatus struct {
	CodeMessage      []ResponseCodeMessage `json:"bCodeMessage"`
	DeveloperMessage string                `json:"developerMessage"`
	StatusCode       int                   `json:"rCode"`
}

// ResponseCodeMessage object
type ResponseCodeMessage struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"errorMessage"`
}

// GreeksResponse object
type GreeksResponse struct {
	Data    GreeksData     `json:"data"`
	Message string         `json:"message"`
	Status  ResponseStatus `json:"status"`
}

// GreeksData object
type GreeksData struct {
	PageTitle string         `json:"pageTitle"`
	Table     GreeksTable    `json:"table"`
	Filters   []GreeksFilter `json:"filters"`
}

// GreeksFilter date of option expiry
type GreeksFilter struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// GreeksTable object
type GreeksTable struct {
	Headers map[string]string `json:"headers"`
	Rows    []GreeksOption    `json:"rows"`
}

// GreeksOption for an option
type GreeksOption struct {
	CallDelta float64 `json:"cDelta"`
	CallGamma float64 `json:"cGamma"`
	CallRho   float64 `json:"cRho"`
	CallTheta float64 `json:"cTheta"`
	CallVega  float64 `json:"cVega"`
	CallIV    float64 `json:"cIV"`
	Strike    float64 `json:"strike"`
	PutDelta  float64 `json:"pDelta"`
	PutGamma  float64 `json:"pGamma"`
	PutRho    float64 `json:"pRho"`
	PutTheta  float64 `json:"pTheta"`
	PutVega   float64 `json:"pVega"`
	PutIV     float64 `json:"pIV"`
	URL       string  `json:"url"`
}

// OptionsResponse object
type OptionsResponse struct {
	Data    OptionsData    `json:"data"`
	Message string         `json:"message"`
	Status  ResponseStatus `json:"status"`
}

// OptionsData object
type OptionsData struct {
	TotalRecords int          `json:"totalRecord"`
	LastTrade    string       `json:"lastTrade"`
	Table        OptionsTable `json:"table"`
	Filters      interface{}  `json:"filterList"` // todo
}

// OptionsTable object
type OptionsTable struct {
	Headers map[string]string `json:"headers"`
	Rows    []Option          `json:"rows"`
}

// Option object
type Option struct {
	ExpiryGroup      string `json:"expirygroup"`
	ExpiryDate       string `json:"expiryDate"`
	CallLast         string `json:"c_Last"`
	CallChange       string `json:"c_Change"`
	CallBid          string `json:"c_Bid"`
	CallAsk          string `json:"c_Ask"`
	CallVolume       string `json:"c_Volume"`
	CallOpeninterest string `json:"c_Openinterest"`
	CallColour       bool   `json:"c_colour"`
	Strike           string `json:"strike"`
	PutLast          string `json:"p_Last"`
	PutChange        string `json:"p_Change"`
	PutBid           string `json:"p_Bid"`
	PutAsk           string `json:"p_Ask"`
	PutVolume        string `json:"p_Volume"`
	PutOpeninterest  string `json:"p_Openinterest"`
	PutColour        bool   `json:"p_colour"`
	URL              string `json:"drillDownURL"`
}

// QuoteResponse object
type QuoteResponse struct {
	Data    QuoteData      `json:"data"`
	Message string         `json:"message"`
	Status  ResponseStatus `json:"status"`
}

// QuoteData object
type QuoteData struct {
	Symbol      string       `json:"symbol"`
	CompanyName string       `json:"companyName"`
	StockType   string       `json:"stockType"`
	Exchange    string       `json:"exchange"`
	PrimaryData PrimaryQuote `json:"primaryData"`
}

// PrimaryQuote object
type PrimaryQuote struct {
	LastSalePrice      string `json:"lastSalePrice"`
	LastTradeTimestamp string `json:"lastTradeTimestamp"`
	IsRealtime         bool   `json:"isRealTime"`
}
