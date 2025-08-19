package main

import (
	"log"

	"github.com/mickaelyoshua/map-recycling-points/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	// UpdateAddresses(config)
}

func UpdateAddresses(config util.Config) {
	log.Printf("Reading addresses JSON file...\n")
	locations, err := ReadJsonFile()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("File Readed.\n\n")

	log.Printf("Getting all coordinates from addresses...\n")
	err = GetAllCoordinates(locations, config.GeocodingApiKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Coordinates retrieved.\n\n")

	log.Printf("Writing to new file with the coordinates...")
	err = WriteJsonFile(locations)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("File saved.\n")
}
