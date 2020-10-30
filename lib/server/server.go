package server

import (
	"feedline/lib/feed"
	"html/template"
	"log"
	"net/http"
)

func Serve(address string) {
	http.Handle("/", indexHandler())
	http.Handle("/subscriptions/", subscriptionsHandler())
	http.Handle("/api/dismiss/", apiDismissHandler())
	log.Fatal(http.ListenAndServe(address, nil))
}

func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("feed")
		var items []feed.Item
		if url == "" {
			items = feed.UnreadItems()
		} else {
			sub, err := feed.FindSubscriptionBySanitizedURL(url)
			if err != nil {
				items = nil
			} else {
				items = sub.UnreadItems()
			}
		}
		t := template.Must(template.ParseFiles("base.gohtml", "feed.gohtml"))
		if err := t.Execute(w, items); err != nil {
			panic(err)
		}
	})
}

func subscriptionsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		subscriptions := feed.Subscriptions()
		t := template.Must(template.ParseFiles("base.gohtml", "subscriptions.gohtml"))
		if err := t.ExecuteTemplate(w, "base.gohtml", subscriptions); err != nil {
			panic(err)
		}
	})
}

func apiDismissHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		feed.MarkAsRead(url)
	})
}
