package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJSON(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONEncode(t *testing.T) {
	logJSON("Hello, World!")
	logJSON(123)
	logJSON(false)
	logJSON([]string{"a", "b", "c"})
	logJSON(map[string]string{"hello": "world"})
}
