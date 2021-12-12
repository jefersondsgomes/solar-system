package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jefersondsgomes/universe-catalog/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
