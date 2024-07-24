package psql

import (
	"go_psql/internal/models"
	"go_psql/internal/config"
	"log"

	_ "github.com/lib/pq"
)

func GetMovies() ([]models.Movie, error) {
	db, err := getConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM " + config.MoviesTableName + " ORDER BY id")
	if err != nil {
		return []models.Movie{}, err
	}
	defer rows.Close()

	movies := make([]models.Movie, 0, 10)

	for rows.Next() {
		m := models.Movie{}
		err = rows.Scan(&m.Id, &m.Title, &m.Director, &m.ReleaseYear, &m.Genre, &m.Duration)
		if err != nil {
			log.Println(err)
			continue
		}
		movies = append(movies, m)
	}
	return movies, nil
}

func InsertMovie(movie models.Movie) error {
	db, err := getConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO "+config.MoviesTableName+"(title, director, release_year, genre, duration) VALUES($1, $2, $3, $4, $5)",
		movie.Title, movie.Director, movie.ReleaseYear, movie.Genre, movie.Duration)

	if err != nil {
		return err
	}
	return nil
}

func UpdateMovie(movie models.Movie) error {
	db, err := getConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE "+config.MoviesTableName+" SET title=$1, director=$2, release_year=$3, genre=$4, duration=$5 WHERE id=$6",
		movie.Title, movie.Director, movie.ReleaseYear, movie.Genre, movie.Duration, movie.Id)
	if err != nil {
		return nil
	}

	return nil
}

func DeleteMovie(id int) error {
	db, err := getConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM "+config.MoviesTableName+" WHERE id=$1", id)
	if err != nil {
		return nil
	}

	return nil
}

func GetMovie(id int) (models.Movie, error) {
	db, err := getConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	m := models.Movie{}
	row := db.QueryRow("SELECT * FROM "+config.MoviesTableName+" WHERE id=$1", id)
	err = row.Scan(&m.Id, &m.Title, &m.Director, &m.ReleaseYear, &m.Genre, &m.Duration)
	if err != nil {
		return models.Movie{}, err
	}

	return m, nil
}
