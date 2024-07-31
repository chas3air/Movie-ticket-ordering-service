package transport

import (
	"go_psql/internal/config"
	"go_psql/internal/database/json"
	"go_psql/internal/models"
	"net/http"
	"strconv"
	"time"
)

func Index_tickets(w http.ResponseWriter, r *http.Request) {
	tickets, err := json.UnmarshalTickets(config.PathJsonFile + "/tickets.json")
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	err = tpl.ExecuteTemplate(w, "tickets_index.html", tickets)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}

func Update_tickets(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id_s := r.Form.Get("id")

	t, err := json.GetTicketByID(config.PathJsonFile+"/tickets.json", id_s)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	if r.Method == http.MethodPost {
		id := r.Form.Get("id")
		movie_title := r.Form.Get("movie_title")
		movie_time_s := r.Form.Get("movie_time")
		viewing_area_s := r.Form.Get("viewing_area")
		email := r.Form.Get("email")

		movie_time_t, _ := time.Parse(`2006-01-02T15:04`, movie_time_s)
		viewing_area_i, _ := strconv.Atoi(viewing_area_s)

		err := json.UpdateTicket(config.PathJsonFile+"/tickets.json", models.Ticket{id, movie_title, movie_time_t, viewing_area_i, email})
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		http.Redirect(w, r, "/tickets", http.StatusSeeOther)
	}

	err = tpl.ExecuteTemplate(w, "tickets_update.html", t)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}

func Delete_tickets(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id_s := r.Form.Get("id")
	err := json.DeleteTicket(config.PathJsonFile + "/tickets.json", id_s)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	http.Redirect(w, r, "/tickets", http.StatusSeeOther)
}