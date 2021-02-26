package utils

// Find value in slice returning index it was found at (or -1)
// and whether or not it was found
func Find(slice []float64, value float64) (int, bool) {
	for idx, item := range slice {
		if item == value {
			return idx, true
		}
	}
	return -1, false
}
