package server

import (
	"feedline/internal/message"
	"feedline/internal/server/api"
	"feedline/pkg/receiver"
	"html/template"
	"log"
	"net/http"
	"path"
)

var webDir string

func Serve(_webDir string, addr string) {
	webDir = _webDir
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/settings/", handleSettings)
	http.HandleFunc("/subscriptions/", handleSubscriptions)

	http.HandleFunc("/api/dismiss/", api.Dismiss)
	http.HandleFunc("/api/feed/", api.Feed)
	http.HandleFunc("/api/refresh/", api.Refresh)
	http.HandleFunc("/api/tidy/", api.Tidy)

	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	var bulletins []receiver.Bulletin
	if url == "" {
		bulletins = receiver.UnreadBulletins()
	} else {
		var err error
		bulletins, err = receiver.AllBulletinsFromSanitizedUrl(url)
		if err != nil {
			panic(err)
		}
	}

	var t *template.Template
	if len(bulletins) != 0 {
		t = template.Must(template.ParseFiles(
			path.Join(webDir, "base.html"),
			path.Join(webDir, "feed.html"),
		))
		if err := t.Execute(w, bulletins); err != nil {
			panic(err)
		}
	} else {
		t = template.Must(template.ParseFiles(
			path.Join(webDir, "base.html"),
			path.Join(webDir, "message.html"),
		))
		if err := t.Execute(w, message.EmptyFeed()); err != nil {
			panic(err)
		}
	}
}

func handleSettings(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	t = template.Must(template.ParseFiles(
		path.Join(webDir, "base.html"),
		path.Join(webDir, "settings.html")))
	if err := t.Execute(w, nil); err != nil {
		panic(err)
	}
}

func handleSubscriptions(w http.ResponseWriter, r *http.Request) {
	channels := receiver.AllChannels()

	var t *template.Template
	if len(channels) != 0 {
		t = template.Must(template.ParseFiles(
			path.Join(webDir, "base.html"),
			path.Join(webDir, "channels.html")))
		if err := t.Execute(w, channels); err != nil {
			panic(err)
		}
	} else {
		t = template.Must(template.ParseFiles(
			path.Join(webDir, "base.html"),
			path.Join(webDir, "message.html"),
		))
		if err := t.Execute(w, message.NoChannels()); err != nil {
			panic(err)
		}
	}
}
