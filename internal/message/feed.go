package message

import "math/rand"

type Message struct {
	Title    string
	Subtitle string
}

func randomize(titles []string, subtitles []string) Message {
	return Message{
		titles[rand.Intn(len(titles))],
		subtitles[rand.Intn(len(subtitles))],
	}
}

func EmptyFeed() Message {
	titles := []string{
		"Nothing to see here",
		"You've caught up",
		"You're all caught up",
	}
	subtitles := []string{
		"Move along, move along",
		"Go about your business",
		"Your pomodoro has failed you",
		"No news is good news",
	}
	return randomize(titles, subtitles)
}

func NoChannels() Message {
	titles := []string{
		"You have no subscriptions",
		"The lights are on, but nobody is home",
	}
	subtitles := []string{
		"This could be a configuration error",
		"Check `subscriptions.opml` for errors",
	}
	return randomize(titles, subtitles)
}
