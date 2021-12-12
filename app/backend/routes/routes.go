package routes

import (
	"github.com/gorilla/mux"
	"github.com/jefersondsgomes/universe-catalog/controllers"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/v1/astro", controllers.Create).Methods("POST")
	router.HandleFunc("/api/v1/astro/{id}", controllers.Get).Methods("GET")
	router.HandleFunc("/api/v1/astros", controllers.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/astro/{id}", controllers.Update).Methods("PUT")
	router.HandleFunc("/api/v1/astro/{id}", controllers.Delete).Methods("DELETE")
}
