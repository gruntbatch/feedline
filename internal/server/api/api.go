package api

import (
	"feedline/pkg/receiver"
	"net/http"
)

func Dismiss(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	receiver.DismissBulletin(url)
}

func Feed(w http.ResponseWriter, r *http.Request) {

}

func Refresh(w http.ResponseWriter, r *http.Request) {
	receiver.Refresh()
}

func Tidy(w http.ResponseWriter, r *http.Request) {
	receiver.TidyDismissed()
}
