package domain

import (
	"errors"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var planetCollection = "planet"

//Planet ...
type Planet struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	Name           string        `bson:"name" json:"name"`
	Climate        string        `bson:"climate" json:"climate"`
	Terrain        string        `bson:"terrain" json:"terrain"`
	AmountOfMovies int           `bson:"amount-of-movies" json:"amount-of-movies"`
}

//PlanetService ...
type PlanetService struct {
	Mongo *mgo.Database
}

//PlanetRepository ...
type PlanetRepository struct {
	Mongo *mgo.Database
}

// Insert ...
func (r *PlanetRepository) Insert(planet *Planet) (*Planet, error) {

	if planet.ID == "" {
		planet.ID = bson.NewObjectId()
	}
	log.Println("Saving: ", planet)
	err := r.Mongo.C(planetCollection).Insert(&planet)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return planet, nil

}

//GetAll ...
func (r *PlanetRepository) GetAll(offset int, limit int) ([]Planet, error) {

	var data []Planet
	r.Mongo.C(planetCollection).Find(bson.M{}).Sort("_id").Limit(limit).Skip(offset).All(&data)

	return data, nil

}

//FindByName ...
func (r *PlanetRepository) FindByName(name string) (Planet, error) {

	planet := Planet{}
	err := r.Mongo.C(planetCollection).Find(bson.M{"name": bson.M{"$regex": bson.RegEx{name + `.*`, "i"}}}).One(&planet)

	if err != nil {
		return planet, err
	}
	return planet, nil
}

//GetByID ...
func (r *PlanetRepository) GetByID(id string) (Planet, error) {

	planet := Planet{}
	if !bson.IsObjectIdHex(id) {
		return planet, errors.New("Planet not found")
	}

	err := r.Mongo.C(planetCollection).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&planet)

	if err != nil {
		return planet, err
	}
	return planet, nil
}

//Delete ...
func (r *PlanetRepository) Delete(id string) error {

	err := r.Mongo.C(planetCollection).Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return err
	}
	return nil
}
