package controllers

import (
	"encoding/json"
	"net/http"
	"solar-system/models"
	"solar-system/repositories"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var astro models.Astro

	err := json.NewDecoder(r.Body).Decode(&astro)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if (models.Astro{}) == astro {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Invalid JSON!")
		return
	}

	newAstro := repositories.Create(astro)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(newAstro)
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	astro := repositories.Get(params["id"])

	if astro.ID == 0 {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(astro)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	astros := repositories.GetAll()
	json.NewEncoder(w).Encode(astros)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var astro models.Astro

	err := json.NewDecoder(r.Body).Decode(&astro)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	params := mux.Vars(r)
	newAstro := repositories.Update(params["id"], astro)

	if newAstro.ID == 0 {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(newAstro)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	repositories.Delete(params["id"])
	w.WriteHeader(204)
}
