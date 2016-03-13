package server

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/asciimoo/privacyscore/checker"
)

var (
	milligramCSS []byte
	milligramURL string = "https://milligram.github.io/css/milligram.min.css"
)

func init() {
	resp, err := http.Get(milligramURL)
	if err != nil {
		log.Fatal("Cannot fetch milligram.css:", err)
	}
	milligramCSS, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Cannot fetch milligram.css:", err)
	}
	initTemplates()
}

func Run(listen *string) error {
	log.Println("listen on", *listen)
	return http.ListenAndServe(*listen, http.HandlerFunc(requestRouter))
}

func requestRouter(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		serveIndexPage(w, r)
	case "/about":
		serveAboutPage(w, r)
	case "/check":
		checkURL(w, r)
	case "/static/milligram.min.css":
		serveMilligramCSS(w, r)
	default:
		http.NotFound(w, r)
	}
}

func serveIndexPage(w http.ResponseWriter, request *http.Request) {
	renderTemplate(w, "index.tpl", nil)
}

func serveAboutPage(w http.ResponseWriter, request *http.Request) {
	renderTemplate(w, "about.tpl", nil)
}

func checkURL(w http.ResponseWriter, request *http.Request) {
	url := request.FormValue("url")
	results, _ := checker.Run(url)
	renderTemplate(w, "result.tpl", results)
}

func serveMilligramCSS(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write(milligramCSS)
}
