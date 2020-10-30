package feed

import (
	"github.com/kennygrant/sanitize"
	"github.com/mmcdole/gofeed"
	"os"
	"sort"
	"time"
)

func parseURL(fp *gofeed.Parser, URL string) []Item {
	rss, err := fp.ParseURL(URL)
	if err != nil {
		return nil
	}

	var items []Item
	for _, item := range rss.Items {
		date := time.Unix(0, 0)
		if item.PublishedParsed != nil {
			date = *item.PublishedParsed
		} else if item.UpdatedParsed != nil {
			date = *item.UpdatedParsed
		}
		items = append(items, Item{item.Title, item.Link, sanitize.BaseName(item.Link), date})
	}

	return items
}

func AllItems() []Item {
	subs := Subscriptions()

	var items []Item

	fp := gofeed.NewParser()
	for _, sub := range subs {
		items = append(items, parseURL(fp, sub.URL)...)
	}

	sort.Slice(items, func(i, j int) bool { return items[i].Date.After(items[j].Date) })

	return items
}

func (s *Subscription) AllItems() []Item {
	items := parseURL(gofeed.NewParser(), s.URL)
	sort.Slice(items, func(i, j int) bool { return items[i].Date.After(items[j].Date) })
	return items
}

func filterReadItems(items []Item) []Item {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	var filtered []Item
	for _, item := range items {
		_, err = os.Stat(home + "/.feedline/read/" + item.SanitizedURL)
		if os.IsNotExist(err) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func UnreadItems() []Item {
	return filterReadItems(AllItems())
}

func (s *Subscription) UnreadItems() []Item {
	return filterReadItems(s.AllItems())
}
