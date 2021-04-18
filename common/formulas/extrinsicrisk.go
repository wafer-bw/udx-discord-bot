package formulas

// GetExtrinsicValue % for provided `share`, `strike`, and `ask`
func GetExtrinsicValue(share float64, strike float64, ask float64) float64 {
	return ((ask - (share - strike)) / share) * 100
}
