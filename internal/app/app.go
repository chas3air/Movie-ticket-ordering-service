package app

import (
	"go_psql/internal/config"
	"go_psql/internal/services"
	"go_psql/internal/transport"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	// sessoin work
	go services.SessionCleaner(config.SessionTable, config.LimitTime)
	go services.ShowSessions(config.SessionTable)

	// router gorilla/max
	r := mux.NewRouter()
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))
	r.HandleFunc("/", transport.Index)
	r.HandleFunc("/login", transport.Login)
	r.HandleFunc("/signup", transport.Signup)
	r.HandleFunc("/profile", transport.ShowProfile)

	// customers admin
	r.HandleFunc("/customers", services.AdminMiddleware(transport.Index_customers))
	r.HandleFunc("/customers/create", services.AdminMiddleware(transport.Create_customers))
	r.HandleFunc("/customers/update", services.AdminMiddleware(transport.Update_customers))
	r.HandleFunc("/customers/delete", services.AdminMiddleware(transport.Delete_customers))

	// movies admin
	r.HandleFunc("/movies/showall", services.AdminMiddleware(transport.Index_movies))
	r.HandleFunc("/movies/create", services.AdminMiddleware(transport.Create_movies))
	r.HandleFunc("/movies/update", services.AdminMiddleware(transport.Update_movies))
	r.HandleFunc("/movies/delete", services.AdminMiddleware(transport.Delete_movies))

	// movies user
	r.HandleFunc("/movies", services.UserMiddleware(transport.ShowMoviesForUser))
	r.HandleFunc("/ordermovie", services.UserMiddleware(transport.OrderTicketToMovie))

	http.ListenAndServe(":8080", r)
}
