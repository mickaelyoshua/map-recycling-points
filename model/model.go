package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"slices"
)

const GeocodingURL = "https://maps.googleapis.com/maps/api/geocode/json"

type Location struct {
	Category  any     `json:"categoria"`
	Name      string  `json:"nome"`
	Address   string  `json:"endereco"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}
type Locations []Location
type ResponseData struct {
	Results []any `json:"results"`
}

func ReadJsonFile(path string) (Locations, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error opening data file: %v\n", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v\n", err)
	}

	var locations []Location
	json.Unmarshal(byteValue, &locations)
	
	return locations, nil
}

func GetCoordinates(client *http.Client, endpoint *url.URL, l *Location, key string) error {
	queryParams := url.Values{}
	queryParams.Set("key", key)
	queryParams.Set("address", l.Address)

	endpoint.RawQuery = queryParams.Encode()

	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return fmt.Errorf("Error creating request: %v\n", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error making request: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var coordLoc ResponseData
	json.Unmarshal(body, &coordLoc)
	
	result := coordLoc.Results[0].(map[string]any)

	l.Address = result["formatted_address"].(string)

	latLng := result["geometry"].(map[string]any)["location"].(map[string]any)

	l.Latitude = latLng["lat"].(float64)
	l.Longitude = latLng["lng"].(float64)

	return nil
}

func GetAllCoordinates(ls Locations, key string) error {
	client := &http.Client{}
	endpoint, err := url.Parse(GeocodingURL)
	if err != nil {
		return fmt.Errorf("Error parsing url: %v\n", err)
	}

	for i := range ls {
		err := GetCoordinates(client, endpoint, &ls[i], key)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteJsonFile(locations []Location, dataPath string) error {
	file, err := json.MarshalIndent(locations, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshalling json: %v\n", err)
	}

	err = os.WriteFile(dataPath, file, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v\n", err)
	}

	return nil
}

func (locations Locations) GetCategories() []string {
	categoriesMap := make(map[string]bool)
	for _, l := range locations {
		
		// category can be string or []string
		if l.Category != "" {
			switch t := l.Category.(type) {
			case string:
				categoriesMap[t] = true
			case []any:
				for _, cat := range t {
					s, ok := cat.(string)
					if ok {
						categoriesMap[s] = true
					}
				}
			}
		}
	}

	var categories []string
	for k := range categoriesMap {
		categories = append(categories, k)
	}
	slices.Sort(categories)
	return categories
}

func (locations Locations) FilterLocations(category string) Locations {
	if category == "Todos" {
		return locations
	}
	
	var filteredLocations Locations
	for _, l := range locations {
		if l.Category == category {
			filteredLocations = append(filteredLocations, l)
		}
	}
	return filteredLocations
}
