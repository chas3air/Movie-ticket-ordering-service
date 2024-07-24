package models

import "fmt"

type Movie struct {
	Id          int
	Title       string
	Director    string
	ReleaseYear int
	Genre       string
	Duration    int
}

func (m Movie) String() string {
	return fmt.Sprintf("id: %d, title: %s, director: %s, release_year: %d, genre: %s, duration: %d",
		m.Id, m.Title, m.Director, m.ReleaseYear, m.Genre, m.Duration)
}
