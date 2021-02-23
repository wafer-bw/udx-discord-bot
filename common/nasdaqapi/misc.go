package nasdaqapi

// ResponseStatus of request
type ResponseStatus struct {
	CodeMessage      string `json:"bCodeMessage"`
	DeveloperMessage string `json:"developerMessage"`
	StatusCode       int    `json:"rCode"`
}
