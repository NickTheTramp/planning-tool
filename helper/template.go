package helper

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Route struct {
	Title    string
	FileName string
}

func RenderTemplate(w http.ResponseWriter, route Route, data interface{}) {
	basePath := filepath.Join("template", "base.html")
	formPath := filepath.Join("template", route.FileName)

	tmpl := template.Must(template.ParseFiles(basePath, formPath))

	err := tmpl.ExecuteTemplate(w, "base.html", route)
	if err != nil {
		return
	}
}
