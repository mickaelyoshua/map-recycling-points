package main

import (
	"encoding/json"
	"io"
	"os"
)

const AddressesDataPath = "data/addresses.json"

type Location struct {
	Category string `json:"categoria"`
	Name     string `json:"nome"`
	Address  string `json:"endereco"`
}

func ReadJsonFile() []Location {
	jsonFile, err := os.Open("data/addresses.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var locations []Location
	json.Unmarshal(byteValue, &locations)
	
	return locations
}
