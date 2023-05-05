package routes

import (
	"github.com/e-phraim/movie-box/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMovieBoxRoutes = func(router *mux.Router) {
	router.HandleFunc("/movie/", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movie/", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{movieId}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{movieId}", controllers.DeleteMovie).Methods("DELETE")
}
