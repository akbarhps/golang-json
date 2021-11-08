package main

import (
	"encoding/json"
	"testing"
)

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		MiddleName: "Doe",
		LastName:   "Smith",
		Age:        25,
		Married:    true,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	t.Log(string(bytes))
}
