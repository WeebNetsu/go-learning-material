package main

import (
	"html/template"
	"log"
	"net/http"
)

type IndexPage struct {
	Title string
	Body  string
}

func indexPageHanler(w http.ResponseWriter, r *http.Request) {
	page := IndexPage{Title: "How to catch fish", Body: "Catching fish is pretty simple, you just buy some worms..."}
	templ, err := template.ParseFiles("index.html")
	if err != nil {
		log.Panic(err)
	}

	templ.Execute(w, page)
}

func main() {
	http.HandleFunc("/", indexPageHanler)
	http.ListenAndServe(":8000", nil)
}
