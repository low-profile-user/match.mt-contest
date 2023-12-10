package controller

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("template/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Index", nil)
}
