package models

import "fmt"

type Movie struct {
    Id          int    `json:"id"`
    Title       string `json:"title"`
    Director    string `json:"director"`
    ReleaseYear int    `json:"releaseYear"`
    Genre       string `json:"genre"`
    Duration    int    `json:"duration"`
}

func (m Movie) String() string {
	return fmt.Sprintf("id: %d, title: %s, director: %s, release_year: %d, genre: %s, duration: %d",
		m.Id, m.Title, m.Director, m.ReleaseYear, m.Genre, m.Duration)
}
