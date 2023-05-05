package models

import (
	"github.com/e-phraim/movie-box/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Movie struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Genre       string `gorm:"genre"`
	Description string `json:"description"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}

// get all
func GetAllMovies() []Movie {
	var Movie []Movie
	db.Find(&Movie)
	return Movie
}


// Create movie
func (m *Movie) CreateMovie() *Movie{
	result := db.Create(&m)
	if result.Error != nil{
		panic(result.Error)
	}
	return m
}

//get by ID
func GetMovieById(Id int64) (*Movie, *gorm.DB){
	var getMovie Movie
	db.Where("ID=?", Id).Find(&getMovie)
	return &getMovie, db
}

//delete movie
func DeleteMovie(ID int64)Movie{
	var movie Movie
	db.Where("ID=?", ID).Delete(&movie)
	return movie
}