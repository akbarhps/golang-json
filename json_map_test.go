package main

import (
	"encoding/json"
	"testing"
)

func TestJSONMap(t *testing.T) {
	data := `{"a":1,"b":2,"c":3}`
	bytes := []byte(data)

	var result map[string]interface{}
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}
