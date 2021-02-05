package handlers

import (
	"net/http"

	"github.com/egamorim/star-wars-planets/pkg/integration"
)

//HandleHealthCheck ...
func (h *Handler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {

	swapi := integration.Swapi{}
	res, err := swapi.GetPlanet("Terra")
	if err != nil {
		h.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	/*
		response := response.HealthResponse{
			Status: "Running.",
		}
	*/
	h.RespondWithJSON(w, http.StatusCreated, res)
}
