package feed

import (
	"errors"
	"feedline/lib/opml"
	"github.com/kennygrant/sanitize"
	"os"
	"time"
)

type Subscription struct {
	Label        string
	URL          string
	SanitizedURL string
}

func FindSubscriptionBySanitizedURL(URL string) (Subscription, error) {
	subs := Subscriptions()

	for _, sub := range subs {
		if sub.SanitizedURL == URL {
			return sub, nil
		}
	}

	return Subscription{}, errors.New("unable to find subscription")
}

var Subscriptions = func() (Subscriptions func() []Subscription) {
	lastModTime := time.Unix(0, 0)
	var subs []Subscription

	return func() []Subscription {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil
		}
		// TODO Load the subscription file from a configured directory
		name := home + "/.feedline/subscriptions.opml"

		stat, err := os.Stat(name)
		if err != nil {
			return nil
		}

		if lastModTime == stat.ModTime() {
			return subs
		}
		lastModTime = stat.ModTime()

		opmlSubs, err := opml.Load(name)
		if err != nil {
			return nil
		}

		subs = flattenSubs(opmlSubs.Body.Outlines)
		return subs
	}
}()

func flattenSubs(outlines []opml.Outline) []Subscription {
	var subs []Subscription
	for _, o := range outlines {
		if o.Type == "rss" {
			subs = append(subs, Subscription{o.Text, o.XMLURL, sanitize.BaseName(o.XMLURL)})
		}
		subs = append(subs, flattenSubs(o.Outlines)...)
	}
	return subs
}

func Follow() {

}

func Unfollow() {

}
