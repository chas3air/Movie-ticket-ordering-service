package json

import (
	"encoding/json"
	"fmt"
	"go_psql/internal/models"
	"os"
)

func ReadMoviesFromFile(filename string) ([]models.Movie, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var movies []models.Movie
	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

// GetMovieByID возвращает фильм по его ID
func GetMovieByID(id int) (models.Movie, error) {
	path, _ := os.Getwd()
	fmt.Println("path: ", path)
	movies, err := ReadMoviesFromFile("internal/database/json/movies.json")
	if err != nil {
		return models.Movie{}, err
	}

	for _, m := range movies {
		if m.Id == id {
			return m, nil
		}
	}

	return models.Movie{}, fmt.Errorf("movie with ID %d not found", id)
}

// CreateMovie добавляет новый фильм в базу данных
func InsertMovie(movie models.Movie) error {
	path, _ := os.Getwd()
	fmt.Println("path: ", path)
	movies, err := ReadMoviesFromFile("internal/database/json/movies.json")
	if err != nil {
		return err
	}

	movies = append(movies, movie)

	return WriteMoviesToFile("internal/database/json/movies.json", movies)
}

// DeleteMovie удаляет фильм из базы данных по его ID
func RemoveMovie(id int) error {
	path, _ := os.Getwd()
	fmt.Println("path: ", path)
	movies, err := ReadMoviesFromFile("internal/database/json/movies.json")
	if err != nil {
		return err
	}

	var updatedMovies []models.Movie
	for _, m := range movies {
		if m.Id != id {
			updatedMovies = append(updatedMovies, m)
		}
	}

	return WriteMoviesToFile("internal/database/json/movies.json", updatedMovies)
}

// WriteMoviesToFile записывает список фильмов в файл
func WriteMoviesToFile(filename string, movies []models.Movie) error {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func UpdateMovie(movie models.Movie) error {
	path, _ := os.Getwd()
	fmt.Println("path: ", path)
	movies, err := ReadMoviesFromFile("internal/database/json/movies.json")
	if err != nil {
		return err
	}

	// Ищем фильм по ID
	for i, m := range movies {
		if m.Id == movie.Id {
			// Обновляем данные фильма
			movies[i] = movie
			return WriteMoviesToFile("internal/database/json/movies.json", movies)
		}
	}

	return fmt.Errorf("movie with ID %d not found", movie.Id)
}