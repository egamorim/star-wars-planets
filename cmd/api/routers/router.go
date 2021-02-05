package routers

import (
	"github.com/egamorim/star-wars-planets/cmd/api/handlers"
	"github.com/gorilla/mux"
)

//Router ...
func Router(h *handlers.Handler) *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/health", h.HandleHealthCheck).Methods("GET")
	router.HandleFunc("/planets", h.HandleInsertNewPlanet).Methods("POST")
	router.HandleFunc("/planets", h.HandleGetAll).Methods("GET")
	router.HandleFunc("/planets/{id}", h.HandleGetByID).Methods("GET")
	router.HandleFunc("/planets/{id}", h.HandleDelete).Methods("DELETE")
	router.HandleFunc("/planets/findByName/{name}", h.HandleFindByName).Methods("GET")
	return router
}
