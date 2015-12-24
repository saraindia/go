package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"fmt"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/home", serveTemplate)
	http.HandleFunc("/robots.txt", handler)
	http.HandleFunc("/error", error_handler)
	http.ListenAndServe(":3000", nil)
}

func error_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ooops... something looks wrong :(")
}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User-agent: *\nDisallow: /")
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "layout.html")

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.Execute(w,tmpl); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}