package router

import (
	"github.com/Tanmoy037/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HeadleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HeadleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HeadleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HeadleFunc("/api/movies/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HeadleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")


	return router
}
