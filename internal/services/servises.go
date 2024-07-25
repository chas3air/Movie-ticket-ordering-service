package services

import (
	"fmt"
	"go_psql/internal/config"
	"go_psql/internal/models"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !AlreadyLoggedIn(w, r, config.CookieName, config.LimitTime, config.UsersTable, config.SessionTable) {
			http.Redirect(w, r, "/", 303)
			return
		}

		u := GetUser(w, r, config.CookieName, config.LimitTime, config.UsersTable, config.SessionTable)
		if u.Role != "admin" {
			http.Error(w, "role is not available", http.StatusNotFound)
			return
		}

		next(w, r)
	}
}

func UserMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !AlreadyLoggedIn(w, r, config.CookieName, config.LimitTime, config.UsersTable, config.SessionTable) {
			http.Redirect(w, r, "/", 303)
			return
		}

		u := GetUser(w, r, config.CookieName, config.LimitTime, config.UsersTable, config.SessionTable)
		if u.Role != "user" && u.Role != "admin" {
			http.Error(w, "role is not available", http.StatusNotFound)
			return
		}

		next(w, r)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, key string, limitTime int, ut map[string]models.Customer, st map[string]models.Session) models.Customer {
	cookie, err := r.Cookie(key)
	if err != nil {
		cookie = &http.Cookie{
			Name:  key,
			Value: uuid.NewV4().String(),
		}
		cookie.MaxAge = limitTime
		http.SetCookie(w, cookie)
	}

	u := models.Customer{}
	s, ok := st[cookie.Name]
	if ok {
		s.LastActivity = time.Now()
		st[cookie.Name] = s
		u = ut[s.UserLogin]
	}

	return u
}

func AlreadyLoggedIn(w http.ResponseWriter, r *http.Request, key string, limitTime int, ut map[string]models.Customer, st map[string]models.Session) bool {
	cookie, err := r.Cookie(key)
	if err != nil {
		return false
	}

	s, ok := st[cookie.Name]
	if ok {
		s.LastActivity = time.Now()
		st[cookie.Value] = s
	}

	_, ok = ut[s.UserLogin]
	cookie.MaxAge = limitTime
	http.SetCookie(w, cookie)
	return ok
}

func SessionCleaner(session_table map[string]models.Session, limitTime int) {
	for {
		for i, v := range session_table {
			if time.Now().Sub(v.LastActivity) > (time.Second * time.Duration(limitTime)) {
				delete(session_table, i)
			}
		}

		time.Sleep(time.Second * 11)
	}
}

func ShowSessions(session_table map[string]models.Session) {
	for {
		fmt.Println("**********")
		for i, v := range session_table {
			fmt.Println(i, " - ", v)
		}

		fmt.Println("**********\n")

		time.Sleep(time.Second * 12)
	}
}
