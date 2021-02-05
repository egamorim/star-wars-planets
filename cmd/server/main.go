package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/egamorim/star-wars-planets/cmd/api/handlers"
	"github.com/egamorim/star-wars-planets/cmd/api/routers"
	"github.com/egamorim/star-wars-planets/pkg/domain"
)

var port string

func init() {
	port = "8000"
}

func main() {

	mongo, err := domain.GetMongoDBDatabase("star-wars")
	if err != nil {
		panic(err)
	}

	planetRepository := domain.PlanetRepository{}
	planetRepository.Mongo = mongo

	h := &handlers.Handler{
		PlanetRepository: &planetRepository,
	}

	routes := routers.Router(h)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      routes,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}

	log.Printf("Star wars is running at port %s\n", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("can't initialize Server", err)
		os.Exit(1)
	}
}
