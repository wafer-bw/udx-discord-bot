package nasdaqapi

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
