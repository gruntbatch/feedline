package main

import (
	"feedline/lib/feed"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", indexHandler())
	http.Handle("/api/dismiss/", apiDismissHandler())
	http.Handle("/api/feed/", apiFeedHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("index.html"))
		t.Execute(w, nil)
	})
}

func apiDismissHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		feed.MarkAsRead(url)
	})
}

func apiFeedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		items := feed.FilterRead(feed.RefreshAll())
		t := template.Must(template.ParseFiles("feed.html"))
		t.Execute(w, items)
	})
}
