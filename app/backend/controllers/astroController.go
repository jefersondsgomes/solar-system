package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../repositories"

	"github.com/gorilla/mux"
)

func GetAstros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	astros := repositories.GetAll()

	if len(astros) == 0 {
		w.WriteHeader(204)
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
		return
	}

	astro := repositories.Get(id)

	if astro.Id != id || astro.Id < 1 {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(astro)
}

func CreateAstro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func UpdateAstro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func DeleteAstro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
