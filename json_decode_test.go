package main

import (
	"encoding/json"
	"testing"
)

func TestJSONDecode(t *testing.T) {
	data := `{"FirstName":"John","MiddleName":"Doe","LastName":"Smith","Age":25,"Married":true}`
	bytes := []byte(data)

	var customer Customer
	err := json.Unmarshal(bytes, &customer)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %s", err)
		panic(err)
	}

	t.Log(customer)
}
