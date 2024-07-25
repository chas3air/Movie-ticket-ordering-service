package transport

import (
	"go_psql/internal/database/psql"
	"go_psql/internal/models"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Index_movies(w http.ResponseWriter, r *http.Request) {
	movies, err := psql.GetMovies()
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	err = tpl.ExecuteTemplate(w, "movies_index.html", movies)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}

func Create_movies(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		title := r.Form.Get("title")
		director := r.Form.Get("director")
		releare_year_s := r.Form.Get("release_year")
		genre := r.Form.Get("genre")
		duration_s := r.Form.Get("duration")

		releare_year_i, err := strconv.Atoi(releare_year_s)
		if err != nil {
			http.Error(w, `<script> alert("year must be number"); </script>`, http.StatusBadRequest)
			return
		}
		duration_i, err := strconv.Atoi(duration_s)
		if err != nil {
			http.Error(w, `<script> alert("duration must be number"); </script>`, http.StatusBadRequest)
			return
		}

		err = psql.InsertMovie(models.Movie{-1, title, director, releare_year_i, genre, duration_i})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/movies/showall", http.StatusBadRequest)
	}

	err := tpl.ExecuteTemplate(w, "movies_create.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}

func Update_movies(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id_s := r.Form.Get("id")
	id_i, _ := strconv.Atoi(id_s)
	m, err := psql.GetMovie(id_i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		title := r.Form.Get("title")
		director := r.Form.Get("director")
		releare_year_s := r.Form.Get("release_year")
		genre := r.Form.Get("genre")
		duration_s := r.Form.Get("duration")

		releare_year_i, err := strconv.Atoi(releare_year_s)
		if err != nil {
			http.Error(w, `<script> alert("year must be number"); </script>`, http.StatusBadRequest)
			return
		}
		duration_i, err := strconv.Atoi(duration_s)
		if err != nil {
			http.Error(w, `<script> alert("duration must be number"); </script>`, http.StatusBadRequest)
			return
		}

		err = psql.UpdateMovie(models.Movie{id_i, title, director, releare_year_i, genre, duration_i})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/movies/showall", http.StatusSeeOther)
	}

	err = tpl.ExecuteTemplate(w, "movies_update.html", m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Delete_movies(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id_s := r.Form.Get("id")
	id_i, _ := strconv.Atoi(id_s)

	err := psql.DeleteMovie(id_i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/movies/showall", http.StatusSeeOther)
}

func ShowMoviesForUser(w http.ResponseWriter, r *http.Request) {
	movies, err := psql.GetMovies()
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	err = tpl.ExecuteTemplate(w, "movies_showall.html", movies)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}

func OrderTicketToMovie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id_s := r.Form.Get("id")
	id_i, _ := strconv.Atoi(id_s)
	m, err := psql.GetMovie(id_i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		email := r.Form.Get("email")

		t := models.Ticket{string(rand.Intn(1000000)), m.Title, time.Now(), 0}

		_ = email
		_ = t

		http.Redirect(w, r, "/movies", 303)
	}

	err = tpl.ExecuteTemplate(w, "movie_ordermovie.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
