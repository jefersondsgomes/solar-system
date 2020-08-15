package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../models"
	"../repositories"

	"github.com/gorilla/mux"
)

func CreateAstro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var astro models.Astro

	err := json.NewDecoder(r.Body).Decode(&astro)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = repositories.Create(astro)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(201)
}

func GetAstros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	astros, err := repositories.GetAll()

	if len(astros) == 0 {
		w.WriteHeader(204)
		return
	} else if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(astros)
}

func GetAstro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	astro, err := repositories.Get(id)

	if astro.Id < 1 {
		w.WriteHeader(204)
		return
	} else if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(astro)
}

func UpdateAstro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 1 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var astro models.Astro
	err = json.NewDecoder(r.Body).Decode(&astro)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = repositories.Update(id, astro)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(200)
}

func DeleteAstro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 1 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = repositories.Delete(id)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(204)
}
