package utils

import (
	"encoding/json"
	"fmt"
)

// PPrint a formatted JSON representation of the provided `obj`
func PPrint(obj interface{}) {
	jdata, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jdata))
}
