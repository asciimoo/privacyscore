package server

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates map[string]*template.Template

func initTemplates() {
	templates = make(map[string]*template.Template)
	templateDir := "server/templates/"
	basePath := templateDir + "base.tpl"
	templateFiles, _ := filepath.Glob(templateDir + "*.tpl")
	for _, f := range templateFiles {
		if f == basePath {
			continue
		}
		templates[filepath.Base(f)] = template.Must(template.ParseFiles(f, basePath))
	}
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		log.Fatal("Template does not exist: ", name)
	}
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal("Template execution error: ", err)
	}
	return err
}