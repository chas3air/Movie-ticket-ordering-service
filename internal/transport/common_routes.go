package transport

import (
	"go_psql/internal/config"
	"go_psql/internal/database/psql"
	"go_psql/internal/models"
	"go_psql/internal/services"
	"go_psql/web"
	"net/http"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

var tpl = web.GetTPL()

func Index(w http.ResponseWriter, r *http.Request) {
	_ = services.GetUser(w, r, config.CookieName, config.LimitTime, config.UsersTable, config.SessionTable)

	err := tpl.ExecuteTemplate(w, "index_choise.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	_ = services.GetUser(w, r, config.CookieName, config.LimitTime, config.UsersTable, config.SessionTable)

	if ok := services.AlreadyLoggedIn(w, r, config.CookieName, config.LimitTime, config.UsersTable, config.SessionTable); ok {
		http.Redirect(w, r, "/movies", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		login := r.Form.Get("login")
		password := r.Form.Get("password")

		c, err := psql.GetPersonWithLoginAndPassword(login, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cookie := &http.Cookie{
			Name:  config.CookieName,
			Value: uuid.NewV4().String(),
		}
		config.SessionTable[cookie.Name] = models.Session{login, time.Now()}
		config.UsersTable[login] = c

		http.Redirect(w, r, "/movies", http.StatusSeeOther)
	}

	err := tpl.ExecuteTemplate(w, "index_login.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	_ = services.GetUser(w, r, config.CookieName, config.LimitTime, config.UsersTable, config.SessionTable)

	if ok := services.AlreadyLoggedIn(w, r, config.CookieName, config.LimitTime, config.UsersTable, config.SessionTable); ok {
		http.Redirect(w, r, "/movies", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		login := r.Form.Get("login")
		password := r.Form.Get("password")
		role := "user"
		name := r.Form.Get("name")
		surname := r.Form.Get("surname")
		age_s := r.Form.Get("age")
		age_i, _ := strconv.Atoi(age_s)

		if c, _ := psql.GetPersonWithLoginAndPassword(login, password); c.Login != "" {
			http.Error(w, "current user already exists", 404)
			return
		}

		new_cust := models.Customer{Login: login, Password: password, Role: role, Name: name, Surname: surname, Age: age_i}
		err := psql.InsertCustomer(new_cust)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cookie := &http.Cookie{
			Name:  config.CookieName,
			Value: uuid.NewV4().String(),
		}

		config.SessionTable[cookie.Name] = models.Session{login, time.Now()}
		config.UsersTable[login] = new_cust

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/movies", 303)
	}

	err := tpl.ExecuteTemplate(w, "index_signup.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}
