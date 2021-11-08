package main

import (
	"encoding/json"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		MiddleName: "Doe",
		LastName:   "Smith",
		Hobbies:    []string{"Skiing", "Snowboarding", "Go"},
		Addresses: []Address{
			{
				Street: "123 Main St",
				City:   "Anytown",
				State:  "CA",
				Zip:    "12345",
			},
			{
				Street: "456 Elm St",
				City:   "Anytown",
				State:  "CA",
				Zip:    "12345",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		t.Error(err)
		panic(err)
	}

	t.Log(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	customer := Customer{}
	err := json.Unmarshal([]byte(`{
		"firstName": "John",
		"middleName": "Doe",
		"lastName": "Smith",
		"hobbies": ["Skiing", "Snowboarding", "Go"],
		"addresses": [
			{
				"street": "123 Main St",
				"city": "Anytown",
				"state": "CA",
				"zip": "12345"
			},
			{
				"street": "456 Elm St",
				"city": "Anytown",
				"state": "CA",
				"zip": "12345"
			}
		]
	}`), &customer)
	if err != nil {
		t.Error(err)
		panic(err)
	}

	t.Log(customer)
	t.Log(customer.Addresses)
}

func TestJSONArrayDecodeDirect(t *testing.T) {
	data := `[
			{
				"street": "123 Main St",
				"city": "Anytown",
				"state": "CA",
				"zip": "12345"
			},
			{
				"street": "456 Elm St",
				"city": "Anytown",
				"state": "CA",
				"zip": "12345"
			}
		]`
	bytes := []byte(data)
	addresses := []Address{}
	err := json.Unmarshal(bytes, &addresses)
	if err != nil {
		t.Error(err)
		panic(err)
	}

	t.Log(addresses)
}
