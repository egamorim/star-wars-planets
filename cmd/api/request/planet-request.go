package request

import (
	"github.com/egamorim/star-wars-planets/pkg/domain"
	"gopkg.in/mgo.v2/bson"
)

//PlanetRequest ...
type PlanetRequest struct {
	ID      bson.ObjectId `json:"id"`
	Name    string        `json:"name" validate:"required"`
	Climate string        `json:"climate" validate:"required"`
	Terrain string        `json:"terrain" validate:"required"`
}

//ToPlanet ...
func (p PlanetRequest) ToPlanet() domain.Planet {

	planet := domain.Planet{
		ID:      p.ID,
		Name:    p.Name,
		Climate: p.Climate,
		Terrain: p.Terrain,
	}
	return planet
}
