package response

import (
	"github.com/egamorim/star-wars-planets/pkg/domain"
)

// PlanetResponse ...
type PlanetResponse struct {
	Planet *domain.Planet
	Links  map[string]string `json:"_links"`
}

// ListPlanetResponse ...
type ListPlanetResponse struct {
	Planets []domain.Planet
	Links   map[string]string `json:"_links"`
}
