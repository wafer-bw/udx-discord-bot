package models

// APIError - Discord API error response object
type APIError struct {
	Message    string  `json:"message"`
	Code       int     `json:"code"`
	RetryAfter float32 `json:"retry_after"`
	Global     bool    `json:"global"`
}
