package handlers

import (
	"net/http"
	"project_movies/db"
	"project_movies/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
	var newMovie models.Movie

	//Bind the JSON payload to the newMovie struct
	if err := c.ShouldBindJSON(&newMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Insert new movie into the database
	_, err := db.DB.Exec("INSERT INTO movies (movie, director, year) VALUES ($1, $2, $3)", newMovie.Movie, newMovie.Director, newMovie.Year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newMovie)
}

func GetMovies(c *gin.Context) {
	var movies []models.Movie

	// Retrieve all movies from the database
	rows, err := db.DB.Query("SELECT id, movie, director, year FROM movies")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.Movie, &movie.Director, &movie.Year); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func MovieById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var movie models.Movie

	// Retrieve the movie with the given ID
	row := db.DB.QueryRow("SELECT  id, movie, director, year FROM movies WHERE id = $1", id)
	if err := row.Scan(&movie.ID, &movie.Movie, &movie.Director, &movie.Year); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func UpdateMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var movieUpdate models.Movie

	if err := c.ShouldBindJSON(&movieUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the movie in the database
	_, err = db.DB.Exec("UPDATE movies SET movie = $1, director = $2, year = $3 WHERE id = $4", movieUpdate.Movie, movieUpdate.Director, movieUpdate.Year, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	movieUpdate.ID = id
	c.JSON(http.StatusOK, movieUpdate)

}
func DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	// Delete the movie from the database
	result, err := db.DB.Exec("DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if any row was affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking rows affected"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No movie found with the given ID"})
		return
	}

	// Send a success response
	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
