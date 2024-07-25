package app

import (
	"go_psql/internal/config"
	"go_psql/internal/services"
	"go_psql/internal/transport"
	"net/http"
)

func Run() {
	go services.SessionCleaner(config.SessionTable, config.LimitTime)
	go services.ShowSessions(config.SessionTable)

	http.HandleFunc("/", transport.Index)
	http.HandleFunc("/login", transport.Login)
	http.HandleFunc("/signup", transport.Signup)

	http.HandleFunc("/customers", services.AdminMiddleware(transport.Index_customers))
	http.HandleFunc("/customers/create", services.AdminMiddleware(transport.Create_customers))
	http.HandleFunc("/customers/update", services.AdminMiddleware(transport.Update_customers))
	http.HandleFunc("/customers/delete", services.AdminMiddleware(transport.Delete_customers))

	http.HandleFunc("/movies", services.UserMiddleware(transport.Index_movies))
	http.HandleFunc("/movies/create", services.AdminMiddleware(transport.Create_movies))
	http.HandleFunc("/movies/update", services.AdminMiddleware(transport.Update_movies))
	http.HandleFunc("/movies/delete", services.AdminMiddleware(transport.Delete_movies))

	http.ListenAndServe(":8080", nil)

}
