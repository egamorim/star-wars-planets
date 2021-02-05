package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/egamorim/star-wars-planets/pkg/domain"
)

//Handler ...
type Handler struct {
	PlanetRepository *domain.PlanetRepository
}

//RespondWithError ...
func (h *Handler) RespondWithError(w http.ResponseWriter, code int, message string) {
	h.RespondWithJSON(w, code, map[string]string{"error": message})
}

//RespondWithJSON ...
func (h *Handler) RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
