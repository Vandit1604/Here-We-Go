package router

import (
	"github.com/gorilla/mux"
	"github.com/vandit1604/mongoapi/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controllers.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/deleteall", controllers.DeleteAllMovies).Methods("DELETE")

	return router
}
