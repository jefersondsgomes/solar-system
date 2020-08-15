package main

import (
	"log"
	"net/http"

	"./controllers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/astros", controllers.GetAstros).Methods("GET")
	r.HandleFunc("/api/v1/astros/{id}", controllers.GetAstro).Methods("GET")
	r.HandleFunc("/api/v1/astros", controllers.CreateAstro).Methods("POST")
	r.HandleFunc("/api/v1/astros/{id}", controllers.UpdateAstro).Methods("PUT")
	r.HandleFunc("/api/v1/astros/{id}", controllers.DeleteAstro).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
