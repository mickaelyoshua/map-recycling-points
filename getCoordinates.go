package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Location struct {
	Category string `json:"categoria"`
	Name     string `json:"nome"`
	Address  string `json:"endereco"`
}

func main() {
	jsonFile, err := os.Open("data/addresses.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var locations []Location
	json.Unmarshal(byteValue, &locations)


}
