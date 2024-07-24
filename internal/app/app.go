package app

import (
	"go_psql/internal/transport"
	"net/http"
)

func Run() {
	http.HandleFunc("/", transport.Index)

	http.HandleFunc("/customers", transport.Index_customers)
	http.HandleFunc("/customers/create", transport.Create_customers)
	http.HandleFunc("/customers/update", transport.Update_customers)
	http.HandleFunc("/customers/delete", transport.Delete_customers)

	http.HandleFunc("/movies", transport.Index_movies)
	http.HandleFunc("/movies/create", transport.Create_movies)
	http.HandleFunc("/movies/update", transport.Update_movies)
	http.HandleFunc("/movies/delete", transport.Delete_movies)

	http.ListenAndServe(":8080", nil)

}
