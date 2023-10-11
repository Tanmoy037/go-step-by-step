package router

import (
	"github.com/Tanmoy037/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HeadleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")

	return router
}