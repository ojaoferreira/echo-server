package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/ojaoferreira/echo-server/services"
)

func GenericHandler(w http.ResponseWriter, r *http.Request) {
	data := services.DataService(r)

	if r.URL.Query().Get("json") == "true" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
		return
	}

	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	tmpl.ExecuteTemplate(w, "Home", data)
}
