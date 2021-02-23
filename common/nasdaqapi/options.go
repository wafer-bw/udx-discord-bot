package nasdaqapi

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
