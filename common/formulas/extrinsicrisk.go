package formulas

// GetExtrinsicRisk % for provided `share`, `strike`, and `ask`
func GetExtrinsicRisk(share float64, strike float64, ask float64) float64 {
	return ((ask - (share - strike)) / share) * 100
}
