package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"runtime"

	"github.com/asciimoo/privacyscore/checker"
)

var (
	milligramCSS []byte
	milligramURL string = "https://milligram.github.io/css/milligram.min.css"
)

var BASE_DIR string

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("No caller information")
	}
	BASE_DIR = path.Dir(filename)
	resp, err := http.Get(milligramURL)
	if err != nil {
		log.Fatal("Cannot fetch milligram.css:", err)
	}
	milligramCSS, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Cannot fetch milligram.css:", err)
	}
	resp.Body.Close()
	initTemplates()
}

func Run(listen *string) error {
	log.Println("listen on", *listen)
	return http.ListenAndServe(*listen, http.HandlerFunc(requestRouter))
}

func requestRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Xss-Protection", "1; mode=block")
	w.Header().Set("X-Content-Type-Options", "nosniff")
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
	results, err := checker.Run(url)
	if err != nil {
		log.Println("[check][error]", url, err)
		renderTemplate(w, "error.tpl", struct {
			Error error
		}{err})
	} else {
		log.Println("[check]", url, results.Score)
		renderTemplate(w, "result.tpl", results)
	}
}

func serveMilligramCSS(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write(milligramCSS)
}
