package utils

import (
	"encoding/json"
)

// FormatJSON returns a indent formatted JSON representation of the provided `obj`
// Panics on any error
func FormatJSON(obj interface{}) string {
	jdata, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(jdata)
}
