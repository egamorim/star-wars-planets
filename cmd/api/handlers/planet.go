package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/go-playground/validator.v9"

	"github.com/egamorim/star-wars-planets/cmd/api/request"
	"github.com/egamorim/star-wars-planets/cmd/api/response"
	"github.com/egamorim/star-wars-planets/pkg/integration"
	"github.com/gorilla/mux"
)

//HandleInsertNewPlanet ...
func (h *Handler) HandleInsertNewPlanet(w http.ResponseWriter, r *http.Request) {
	v := validator.New()

	req := new(request.PlanetRequest)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		log.Println(err.Error())
		h.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := v.Struct(req)
	if err != nil {
		h.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	swapi := integration.Swapi{}
	swapiPlanet, err := swapi.GetPlanet(req.Name)
	if err != nil {
		h.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	planet := req.ToPlanet()
	planet.AmountOfMovies = len(swapiPlanet.Films)
	p, err := h.PlanetRepository.Insert(&planet)

	if err != nil {
		log.Println("Error:", err.Error())
		h.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := response.PlanetResponse{Planet: p, Links: make(map[string]string)}
	response.Links["self"] = fmt.Sprintf("http://%s%s/%s", r.Host, r.URL.Path, string(p.ID.Hex()))

	h.RespondWithJSON(w, http.StatusCreated, response)
}

//HandleGetAll ...
func (h *Handler) HandleGetAll(w http.ResponseWriter, r *http.Request) {

	l, _ := r.URL.Query()["limit"]
	o, _ := r.URL.Query()["offset"]

	limit := 5
	offset := 0

	if l != nil {
		limit, _ = strconv.Atoi(l[0])
	}
	if o != nil {
		offset, _ = strconv.Atoi(o[0])
	}

	planets, err := h.PlanetRepository.GetAll(offset, limit)
	if err != nil {
		log.Println("Error:", err.Error())
		h.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	nextOffset := offset + limit
	previousOffset := offset - limit

	response := response.ListPlanetResponse{Planets: planets, Links: make(map[string]string)}
	response.Links["next"] = fmt.Sprintf("http://%s%s?offset=%d&limit=%d", r.Host, r.URL.Path, nextOffset, limit)

	if previousOffset >= 0 {
		response.Links["prev"] = fmt.Sprintf("http://%s%s?offset=%d&limit=%d", r.Host, r.URL.Path, previousOffset, limit)
	}

	h.RespondWithJSON(w, http.StatusOK, response)
}

//HandleFindByName ...
func (h *Handler) HandleFindByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name, _ := vars["name"]

	p, err := h.PlanetRepository.FindByName(name)
	if err != nil {
		log.Println("Error:", err.Error())
		h.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	h.RespondWithJSON(w, http.StatusOK, p)
}

//HandleGetByID ...
func (h *Handler) HandleGetByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := vars["id"]

	p, err := h.PlanetRepository.GetByID(id)
	if err != nil {
		log.Println("Error:", err.Error())
		h.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	h.RespondWithJSON(w, http.StatusOK, p)
}

//HandleDelete ...
func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := vars["id"]

	err := h.PlanetRepository.Delete(id)
	if err != nil {
		log.Println("Error:", err.Error())
		h.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	h.RespondWithJSON(w, http.StatusOK, nil)
}
