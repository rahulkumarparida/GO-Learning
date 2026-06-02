package router

import (
	"mongosetup/controller"

	"github.com/gorilla/mux"
)




func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", controller.MarkMovieWatched).Methods("PATCH")
	router.HandleFunc("/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/movies", controller.DeleteMovies).Methods("DELETE")

	return router
}