package server

import (
	"feedline/internal/server/api"
	"feedline/internal/server/page"
	"log"
	"net/http"
)

var webDir string

func Serve(webDir string, addr string) {
	page.WebDir = webDir

	http.HandleFunc("/api/dismiss/", api.Dismiss)
	http.HandleFunc("/api/feed/", api.Feed)
	http.HandleFunc("/api/refresh/", api.Refresh)
	http.HandleFunc("/api/tidy/", api.Tidy)

	http.HandleFunc("/", page.Index)
	http.HandleFunc("/settings/", page.Settings)
	http.HandleFunc("/subscriptions/", page.Subscriptions)

	log.Fatal(http.ListenAndServe(addr, nil))
}
