package main

import (
	"project_movies/db"
	"project_movies/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.SetupDatabase()

	//define routes

	router.POST("/movies", handlers.CreateMovie)
	router.GET("/allmovies", handlers.GetMovies)
	router.GET("/movies/id", handlers.MovieById)
	router.PUT("/movies/id", handlers.UpdateMovie)
	router.DELETE("/movies/id", handlers.DeleteMovie)

	router.Run(":8080")
}
