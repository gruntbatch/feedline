package feed

import (
	"github.com/kennygrant/sanitize"
	"github.com/mmcdole/gofeed"
	"sort"
	"time"
)

func RefreshAll() []Item {
	feeds := All()

	var items []Item

	for _, feed := range feeds {
		fp := gofeed.NewParser()
		rss, err := fp.ParseURL(feed.URL)
		if err != nil {
			panic(err)
		}

		for _, item := range rss.Items {
			date := time.Unix(0, 0)
			if item.PublishedParsed != nil {
				date = *item.PublishedParsed
			} else if item.UpdatedParsed != nil {
				date = *item.UpdatedParsed
			}
			items = append(items, Item{item.Title, item.Link, sanitize.BaseName(item.Link), date})
		}
	}

	sort.Slice(items, func(i, j int) bool { return items[i].Date.Before(items[j].Date) })

	return items
}
