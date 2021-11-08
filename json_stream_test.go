package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestJSONStreamDecoder(t *testing.T) {
	reader, _ := os.Open("./resources/input.json")
	decoder := json.NewDecoder(reader)

	product := &Product{}
	err := decoder.Decode(product)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
		panic(err)
	}

	t.Log(product)
}

func TestJSONStreamEncoder(t *testing.T) {
	product := &Product{
		Id:       "MBP-1",
		Name:     "MacBook Pro",
		Price:    123123,
		ImageURL: "mekbuk.com",
	}

	writer, _ := os.Create("./resources/output.json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(product)

	if err != nil {
		t.Errorf("Error encoding JSON: %s", err)
		panic(err)
	}
}
