package feed

import (
	"feedline/lib/opml"
	"github.com/kennygrant/sanitize"
	"os"
	"time"
)

func All() []Item {
	// TODO Load the subscription file only if it has changed
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	// TODO Load the subscription file from a configured directory
	opmlSubs, err := opml.Load(homeDir + "/.feedline/subscriptions.opml")
	if err != nil {
		return nil
	}

	feeds := flattenFeeds(opmlSubs.Body.Outlines)
	return feeds
}

func flattenFeeds(outlines []opml.Outline) []Item {
	var feeds []Item
	for _, o := range outlines {
		if o.Type == "rss" {
			feeds = append(feeds, Item{o.Text, o.XMLURL, sanitize.BaseName(o.XMLURL), time.Unix(0, 0)})
		}
		feeds = append(feeds, flattenFeeds(o.Outlines)...)
	}
	return feeds
}

func Follow() {

}

func Unfollow() {

}
