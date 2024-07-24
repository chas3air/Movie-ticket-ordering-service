package transport

import (
	"go_psql/web"
	"net/http"
)

var tpl = web.GetTPL()

func Index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index_choise.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}
