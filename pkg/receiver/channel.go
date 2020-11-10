package receiver

import (
	"feedline/pkg/opml"
	"github.com/kennygrant/sanitize"
	"github.com/mmcdole/gofeed"
	"os"
	"sync"
	"time"
)

type ItemInfo struct {
	Label        string
	URL          string
	SanitizedURL string
}

type Channel struct {
	ItemInfo
}

func AllChannels() []Channel {
	linked := linkedChannels()
	var channels []Channel

	for linked != nil {
		channels = append(channels, Channel{
			linked.Info,
		})
		linked = linked.next
	}

	return channels
}

type LinkedChannel struct {
	mu        sync.RWMutex
	next      *LinkedChannel
	bulletins []Bulletin
	Info      ItemInfo
}

var linkedChannels = func() (linkedChannels func() *LinkedChannel) {
	var head *LinkedChannel = nil

	mu := sync.RWMutex{}
	last := time.Unix(0, 0)

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return func() *LinkedChannel {
		// TODO Load the subscription file from a configured directory
		name := home + "/.feedline/subscriptions.opml"

		stat, err := os.Stat(name)
		if err != nil {
			panic(err)
		}

		mu.Lock()
		defer mu.Unlock()

		if last.Equal(stat.ModTime()) {
			return head
		} else {
			last = stat.ModTime()
			opml, err := opml.Load(name)
			if err != nil {
				panic(err)
			}

			head = linkOutline(opml.Body.Outlines, nil)
			return head
		}
	}
}()

func linkOutline(outlines []opml.Outline, head *LinkedChannel) *LinkedChannel {
	for _, o := range outlines {
		if o.Type == "rss" {
			head = &LinkedChannel{
				sync.RWMutex{},
				head,
				nil,
				ItemInfo{o.Text, o.XMLURL, sanitize.BaseName(o.XMLURL)},
			}
		}
		head = linkOutline(o.Outlines, head)
	}
	return head
}

func (c *LinkedChannel) refresh() {
	c.mu.RLock()
	url := c.Info.URL
	c.mu.RUnlock()

	fp := gofeed.NewParser()
	rss, err := fp.ParseURL(url)
	if err != nil {
		panic(err)
	}

	var bulletins []Bulletin
	for _, item := range rss.Items {
		date := time.Unix(0, 0)
		if item.PublishedParsed != nil {
			date = *item.PublishedParsed
		} else if item.UpdatedParsed != nil {
			date = *item.UpdatedParsed
		}
		bulletins = append(bulletins, Bulletin{
			ItemInfo{item.Title, item.Link, sanitize.BaseName(item.Link)},
			date,
			c.Info,
		})
	}

	c.mu.Lock()
	c.bulletins = bulletins
	c.mu.Unlock()
}
