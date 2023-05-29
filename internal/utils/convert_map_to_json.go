package utils

import (
	"encoding/json"
)

// Map2JSON returns the JSON encoding by given map.
func Map2JSON(m map[string]interface{}) string {
	b, err := json.Marshal(m)
	if err != nil {
		return `{"logging-error":"error while marshalling"}`
	}

	return string(b)
}
