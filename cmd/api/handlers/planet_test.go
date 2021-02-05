package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/egamorim/star-wars-planets/cmd/api/response"

	"github.com/egamorim/star-wars-planets/cmd/api/handlers"
	"github.com/egamorim/star-wars-planets/cmd/api/routers"
	"github.com/egamorim/star-wars-planets/pkg/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	planetTestDatabase    = "star-wars-test"
	apiPlanetResourcePath = "/planets"
)

func TestHandleInsertNewPlanet_Success(t *testing.T) {

	mongo, _ := domain.GetMongoDBDatabase(planetTestDatabase)
	defer dropTestDatabase(mongo)

	payload := []byte(`{"name": "Tatooine", "terrain" : "mais um", "climate" : "gelado"}`)
	req, _ := http.NewRequest("POST", apiPlanetResourcePath, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	res := executeRequest(req, mongo)

	checkResponseCode(t, http.StatusCreated, res.Code)
}

func TestHandleInsertNewPlanet_PlanetNotFoundSwapi(t *testing.T) {

	mongo, _ := domain.GetMongoDBDatabase(planetTestDatabase)
	defer dropTestDatabase(mongo)

	payload := []byte(`{"name": "Terra", "terrain" : "mais um", "climate" : "gelado"}`)
	req, _ := http.NewRequest("POST", apiPlanetResourcePath, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	res := executeRequest(req, mongo)

	checkResponseCode(t, http.StatusNotFound, res.Code)
}

func TestHandleGetAll(t *testing.T) {

	mongo, _ := domain.GetMongoDBDatabase(planetTestDatabase)
	defer dropTestDatabase(mongo)

	loadPlanetData(mongo)

	req, _ := http.NewRequest("GET", apiPlanetResourcePath, nil)

	res := executeRequest(req, mongo)

	var l response.ListPlanetResponse
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&l)

	checkResponseCode(t, http.StatusOK, res.Code)

	if len(l.Planets) != 4 {
		t.Errorf("Expected 4 but got %d", len(l.Planets))
	}
}

func TestHandleFindByName_Success(t *testing.T) {

	mongo, _ := domain.GetMongoDBDatabase(planetTestDatabase)
	defer dropTestDatabase(mongo)

	loadPlanetData(mongo)

	req, _ := http.NewRequest("GET", apiPlanetResourcePath+"/findByName/Tat", nil)

	res := executeRequest(req, mongo)

	var p domain.Planet
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&p)

	checkResponseCode(t, http.StatusOK, res.Code)

	if p.Name != "Tatooine" {
		t.Errorf("Expected 'Tatooine' but got %s", p.Name)
	}
}

func TestHandleFindByName_PlanetNotFound(t *testing.T) {

	mongo, _ := domain.GetMongoDBDatabase(planetTestDatabase)
	defer dropTestDatabase(mongo)

	loadPlanetData(mongo)

	req, _ := http.NewRequest("GET", apiPlanetResourcePath+"/findByName/Tataaa", nil)
	res := executeRequest(req, mongo)

	checkResponseCode(t, http.StatusNotFound, res.Code)

}

func TestHandleGetByID_Success(t *testing.T) {

	mongo, _ := domain.GetMongoDBDatabase(planetTestDatabase)
	defer dropTestDatabase(mongo)
	loadPlanetData(mongo)

	req, _ := http.NewRequest("GET", apiPlanetResourcePath+"/5ba005525f4f723340e0eb83", nil)
	res := executeRequest(req, mongo)

	checkResponseCode(t, http.StatusOK, res.Code)

	var p domain.Planet
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&p)

	checkResponseCode(t, http.StatusOK, res.Code)

	if p.Name != "Yavin IV" {
		t.Errorf("Expected 'Tatooine' but got %s", p.Name)
	}

}

func TestHandleGetByID_NotFound(t *testing.T) {

	mongo, _ := domain.GetMongoDBDatabase(planetTestDatabase)
	defer dropTestDatabase(mongo)
	loadPlanetData(mongo)

	req, _ := http.NewRequest("GET", apiPlanetResourcePath+"/5ba005525f4f723340e0eb83TEST", nil)
	res := executeRequest(req, mongo)
	checkResponseCode(t, http.StatusNotFound, res.Code)

}

func TestHandleDelete(t *testing.T) {

	mongo, _ := domain.GetMongoDBDatabase(planetTestDatabase)
	defer dropTestDatabase(mongo)

	loadPlanetData(mongo)

	req, _ := http.NewRequest("GET", apiPlanetResourcePath, nil)
	res := executeRequest(req, mongo)

	var l response.ListPlanetResponse
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&l)

	checkResponseCode(t, http.StatusOK, res.Code)

	if len(l.Planets) != 4 {
		t.Errorf("Expected 4 but got %d", len(l.Planets))
	}

	req, _ = http.NewRequest("DELETE", apiPlanetResourcePath+"/5ba005525f4f723340e0eb83", nil)
	res = executeRequest(req, mongo)
	checkResponseCode(t, http.StatusOK, res.Code)

	req, _ = http.NewRequest("GET", apiPlanetResourcePath, nil)
	res = executeRequest(req, mongo)

	decoder = json.NewDecoder(res.Body)
	decoder.Decode(&l)

	if len(l.Planets) != 3 {
		t.Errorf("Expected 3 but got %d", len(l.Planets))
	}

}

func executeRequest(req *http.Request, mongo *mgo.Database) *httptest.ResponseRecorder {

	planetRepository := domain.PlanetRepository{}
	planetRepository.Mongo = mongo

	rr := httptest.NewRecorder()
	h := &handlers.Handler{
		PlanetRepository: &planetRepository,
	}

	routers.Router(h).ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func loadPlanetData(mongo *mgo.Database) {

	collection := mongo.C("planet")

	collection.Insert(&domain.Planet{
		ID:      bson.ObjectIdHex("5ba005525f4f723340e0eb81"),
		Name:    "Tatooine",
		Climate: "arid",
		Terrain: "desert",
	})

	collection.Insert(&domain.Planet{
		ID:      bson.ObjectIdHex("5ba005525f4f723340e0eb82"),
		Name:    "Alderaan",
		Climate: "temperate",
		Terrain: "grasslands, mountains",
	})

	collection.Insert(&domain.Planet{
		ID:      bson.ObjectIdHex("5ba005525f4f723340e0eb83"),
		Name:    "Yavin IV",
		Climate: "temperate, tropical",
		Terrain: "jungle, rainforests",
	})

	collection.Insert(&domain.Planet{
		ID:      bson.ObjectIdHex("5ba005525f4f723340e0eb84"),
		Name:    "Hoth",
		Climate: "frozen",
		Terrain: "tundra, ice caves, mountain ranges",
	})

}

func dropTestDatabase(mongo *mgo.Database) {
	mongo.DropDatabase()
}
