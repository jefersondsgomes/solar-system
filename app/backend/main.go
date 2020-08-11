package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/astros", nil).Methods("GET")
	r.HandleFunc("/api/v1/astros/{id}", nil).Methods("GET")
	r.HandleFunc("/api/v1/astros", nil).Methods("POST")
	r.HandleFunc("/api/v1/astros/{id}", nil).Methods("PUT")
	r.HandleFunc("/api/v1/astros/{id}", nil).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
