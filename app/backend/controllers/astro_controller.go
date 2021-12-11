package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jefersondsgomes/solar-system-catalog/helpers"
	"github.com/jefersondsgomes/solar-system-catalog/models"
	"github.com/jefersondsgomes/solar-system-catalog/repositories"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var astro models.Astro
	if err := json.NewDecoder(r.Body).Decode(&astro); err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(helpers.GenerateErrorResponse(400, err.Error()))
		return
	}

	newAstro, err := repositories.Create(astro)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(helpers.GenerateErrorResponse(500, err.Error()))
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(newAstro)
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 0, 0)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(helpers.GenerateErrorResponse(500, err.Error()))
		return
	}

	astro := models.Astro{ID: id}
	var defaultAstro = astro
	astro, err = repositories.Get(astro)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(helpers.GenerateErrorResponse(500, err.Error()))
		return
	}

	if astro == defaultAstro {
		json.NewEncoder(w).Encode(helpers.Empty{})
		return
	} else {
		json.NewEncoder(w).Encode(astro)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	astros, err := repositories.GetAll()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(helpers.GenerateErrorResponse(500, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(astros)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var astro models.Astro
	if err := json.NewDecoder(r.Body).Decode(&astro); err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(helpers.GenerateErrorResponse(400, err.Error()))
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 0, 0)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(helpers.GenerateErrorResponse(500, err.Error()))
		return
	}

	astro.ID = id
	astro.PhysicalData.ID = id
	astro.PhysicalData.AstroID = id
	var defaultAstro = astro
	astro, err = repositories.Update(astro)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(helpers.GenerateErrorResponse(500, err.Error()))
		return
	}

	if astro == defaultAstro {
		w.WriteHeader(201)
		return
	}

	json.NewEncoder(w).Encode(astro)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 0, 0)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(helpers.GenerateErrorResponse(500, err.Error()))
		return
	}

	astro := models.Astro{ID: id}
	if err = repositories.Delete(astro); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(helpers.GenerateErrorResponse(500, err.Error()))
		return
	}

	w.WriteHeader(204)
}
