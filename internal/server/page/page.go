package page

import (
	"feedline/internal/message"
	"feedline/pkg/receiver"
	"net/http"
	"path"
	"text/template"
)

var WebDir string

func Index(w http.ResponseWriter, r *http.Request) {
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
			path.Join(WebDir, "base.html"),
			path.Join(WebDir, "feed.html"),
		))
		if err := t.Execute(w, bulletins); err != nil {
			panic(err)
		}
	} else {
		t = template.Must(template.ParseFiles(
			path.Join(WebDir, "base.html"),
			path.Join(WebDir, "message.html"),
		))
		if err := t.Execute(w, message.EmptyFeed()); err != nil {
			panic(err)
		}
	}
}

func Settings(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	t = template.Must(template.ParseFiles(
		path.Join(WebDir, "base.html"),
		path.Join(WebDir, "settings.html")))
	if err := t.Execute(w, nil); err != nil {
		panic(err)
	}
}

func Subscriptions(w http.ResponseWriter, r *http.Request) {
	channels := receiver.AllChannels()

	var t *template.Template
	if len(channels) != 0 {
		t = template.Must(template.ParseFiles(
			path.Join(WebDir, "base.html"),
			path.Join(WebDir, "channels.html")))
		if err := t.Execute(w, channels); err != nil {
			panic(err)
		}
	} else {
		t = template.Must(template.ParseFiles(
			path.Join(WebDir, "base.html"),
			path.Join(WebDir, "message.html"),
		))
		if err := t.Execute(w, message.NoChannels()); err != nil {
			panic(err)
		}
	}
}

func Static(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	http.ServeFile(w, r, path.Join(WebDir, r.URL.Path[1:]))
}
