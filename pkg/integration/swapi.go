package integration

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

//BaseURL ...
const BaseURL = "http://swapi.dev/api/planets?search="

var client = &http.Client{}

//Swapi ...
type Swapi struct {
}

// SwapiPlanet ...
type SwapiPlanet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	Films          []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
}

//SwapiPlanetResponse ...
type SwapiPlanetResponse struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []SwapiPlanet `json:"results"`
}

//GetPlanet ...
func (s Swapi) GetPlanet(name string) (SwapiPlanet, error) {

	url := BaseURL + name
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}

	defer res.Body.Close()

	externalResponse := SwapiPlanetResponse{}
	decodeErr := json.NewDecoder(res.Body).Decode(&externalResponse)
	if decodeErr != nil {
		log.Printf("Error parsing response: %s\n", decodeErr)
	}

	if len(externalResponse.Results) == 0 {
		return SwapiPlanet{}, errors.New("Planet not found in Swapi service")
	}
	return externalResponse.Results[0], nil
}
