package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/e-phraim/movie-box/pkg/models"
	"github.com/e-phraim/movie-box/pkg/utils"
	"github.com/gorilla/mux"
)

var NewMovie models.Movie

// get all
func GetMovie(w http.ResponseWriter, r *http.Request) {
	newMovies := models.GetAllMovies()
	res, _ := json.Marshal(newMovies)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// create a movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	CreateMovie := &models.Movie{}
	utils.ParseBody(r, CreateMovie)
	m := CreateMovie.CreateMovie()
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// get by id
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing Record")
	}
	movieDetails, _ := models.GetMovieById(ID)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// update movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var updateMovie = &models.Movie{}
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing Record")
	}
	movieDetails, db := models.GetMovieById(ID)
	if updateMovie.Name != ""{
		updateMovie.Name = updateMovie.Name
	}
	if updateMovie.Genre != ""{
		updateMovie.Genre = updateMovie.Genre
	}
	if updateMovie.Description != ""{
		updateMovie.Description = updateMovie.Description
	}
	db.Save(&movieDetails)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//delete movie
func DeleteMovie(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId,0,0)
	if err != nil{
		fmt.Println("Error While Parsing Record")
	}
	movie := models.DeleteMovie(ID)
	res, _ := json.Marshal(movie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}