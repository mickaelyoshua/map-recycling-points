package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua/map-recycling-points/controller"
	"github.com/mickaelyoshua/map-recycling-points/model"
	"github.com/mickaelyoshua/map-recycling-points/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	// UpdateAddresses(config)

	locations, err := model.ReadJsonFile(config.FinalDataPath)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Static("/static", "./static")

	router.GET("/", controller.Index(locations))

	router.Run(":" + config.ServerPort)
}

func UpdateAddresses(config util.Config) {
	log.Printf("Reading addresses JSON file...\n")
	locations, err := model.ReadJsonFile(config.AddressesDataPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("File Readed.\n\n")

	log.Printf("Getting all coordinates from addresses...\n")
	err = model.GetAllCoordinates(locations, config.GeocodingApiKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Coordinates retrieved.\n\n")

	log.Printf("Writing to new file with the coordinates...")
	err = model.WriteJsonFile(locations, config.FinalDataPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("File saved.\n")
}
