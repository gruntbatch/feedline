package server

import (
	"feedline/lib/receiver"
	"html/template"
	"log"
	"net/http"
)

func Serve(address string) {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/subscriptions/", handleSubscriptions)
	http.HandleFunc("/api/dismiss/", handleDismiss)
	http.HandleFunc("/api/refresh/", handleRefresh)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	var bulletins []receiver.Bulletin
	if url == "" {
		bulletins = receiver.UnreadBulletins()
	} else {
		var err error
		bulletins, err = receiver.UnreadBulletinsFromSanitizedURL(url)
		if err != nil {
			panic(err)
		}
	}
	t := template.Must(template.ParseFiles("base.gohtml", "feed.gohtml"))
	if err := t.Execute(w, bulletins); err != nil {
		panic(err)
	}
}

func handleSubscriptions(w http.ResponseWriter, r *http.Request) {
	channels := receiver.AllChannels()
	t := template.Must(template.ParseFiles("base.gohtml", "channels.gohtml"))
	if err := t.ExecuteTemplate(w, "base.gohtml", channels); err != nil {
		panic(err)
	}
}

func handleDismiss(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	receiver.DismissBulletin(url)
}

func handleRefresh(w http.ResponseWriter, r *http.Request) {
	receiver.Refresh()
}
