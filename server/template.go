package server

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/asciimoo/privacyscore/utils"
)

var templates map[string]*template.Template

var funcMap = template.FuncMap{
	"GetScoreName": utils.GetScoreName,
	"statHeight":   statHeight,
}

func initTemplates() {
	templates = make(map[string]*template.Template)
	templateDir := BASE_DIR + "/templates/"
	basePath := templateDir + "base.tpl"
	templateFiles, _ := filepath.Glob(templateDir + "*.tpl")
	for _, f := range templateFiles {
		if f == basePath {
			continue
		}
		templates[filepath.Base(f)] = template.Must(template.New("").Funcs(funcMap).ParseFiles(f, basePath))
	}
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		log.Fatal("Template does not exist: ", name)
	}
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println("Template execution error: ", err)
	}
	return err
}

func statHeight(a, b uint) uint {
	if b == 0 {
		return 0
	}
	return a * 100 / b
}
