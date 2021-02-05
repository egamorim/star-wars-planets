package handlers

import (
	"net/http"

	"github.com/egamorim/star-wars-planets/cmd/api/response"
)

//HandleHealthCheck ...
func (h *Handler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {

	response := response.HealthResponse{
		Status: "Running.",
	}
	h.RespondWithJSON(w, http.StatusCreated, response)
}
