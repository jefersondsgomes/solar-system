package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jefersondsgomes/universe-catalog/entities"
	"github.com/jefersondsgomes/universe-catalog/repositories"
	"github.com/jefersondsgomes/universe-catalog/utils"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var astro entities.Astro
	if err := json.NewDecoder(r.Body).Decode(&astro); err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(utils.GenerateErrorResponse(400, err.Error()))
		return
	}

	newAstro, err := repositories.Create(astro)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(utils.GenerateErrorResponse(500, err.Error()))
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
		json.NewEncoder(w).Encode(utils.GenerateErrorResponse(500, err.Error()))
		return
	}

	astro := entities.Astro{ID: id}
	var defaultAstro = astro
	astro, err = repositories.Get(astro)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(utils.GenerateErrorResponse(500, err.Error()))
		return
	}

	if astro == defaultAstro {
		json.NewEncoder(w).Encode(utils.Empty{})
		return
	} else {
		json.NewEncoder(w).Encode(astro)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pagination := utils.GeneratePagination(r)
	astros, err := repositories.GetAll(pagination)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(utils.GenerateErrorResponse(500, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(astros)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var astro entities.Astro
	if err := json.NewDecoder(r.Body).Decode(&astro); err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(utils.GenerateErrorResponse(400, err.Error()))
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 0, 0)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(utils.GenerateErrorResponse(500, err.Error()))
		return
	}

	astro.ID = id
	astro, err = repositories.Update(astro)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(utils.GenerateErrorResponse(500, err.Error()))
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
		json.NewEncoder(w).Encode(utils.GenerateErrorResponse(500, err.Error()))
		return
	}

	astro := entities.Astro{ID: id}
	if err = repositories.Delete(astro); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(utils.GenerateErrorResponse(500, err.Error()))
		return
	}

	w.WriteHeader(204)
}
