package transport

import (
	"go_psql/internal/database/psql"
	"go_psql/internal/models"
	"net/http"
	"strconv"
)

func Index_customers(w http.ResponseWriter, r *http.Request) {
	custs, err := psql.GetPeople()
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	err = tpl.ExecuteTemplate(w, "customers_index.html", custs)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}

func Create_customers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		login := r.Form.Get("login")
		password := r.Form.Get("password")
		name := r.Form.Get("name")
		surname := r.Form.Get("surname")
		age_i, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			http.Error(w, `<script> alert("age must be number"); </script>`, http.StatusBadRequest)
			return
		}

		err = psql.InsertCustomer(models.Customer{Login: login, Password: password, Name: name, Surname: surname, Age: age_i})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/customers", http.StatusSeeOther)
	}

	err := tpl.ExecuteTemplate(w, "customers_create.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Update_customers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id_s := r.Form.Get("id")
	id_i, _ := strconv.Atoi(id_s)
	c, err := psql.GetPerson(id_i)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	if r.Method == http.MethodPost {
		login := r.Form.Get("login")
		password := r.Form.Get("password")
		name := r.Form.Get("name")
		surname := r.Form.Get("surname")
		age_s := r.Form.Get("age")
		age_i, _ := strconv.Atoi(age_s)

		err := psql.UpdateCustomer(models.Customer{id_i, login, password, name, surname, age_i})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/customers", http.StatusSeeOther)
	}

	err = tpl.ExecuteTemplate(w, "customers_update.html", c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Delete_customers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id_s := r.Form.Get("id")
	id_i, _ := strconv.Atoi(id_s)

	err := psql.DeleteCustomer(id_i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/customers", http.StatusSeeOther)
}
