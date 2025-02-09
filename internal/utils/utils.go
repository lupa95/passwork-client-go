package utils

import (
	"encoding/json"
	"fmt"
)

// ParseJSONResponse unmarshals JSON data into a given struct type.
func ParseJSONResponse[T any](data []byte) (T, error) {
	var responseObject T
	err := json.Unmarshal(data, &responseObject)
	if err != nil {
		return responseObject, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return responseObject, nil
}
