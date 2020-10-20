package feed

import (
	"github.com/kennygrant/sanitize"
	"os"
	"time"
)

func MarkAsRead(URL string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	name := homeDir + "/.feedline/read/" + sanitize.BaseName(URL)

	_, err = os.Stat(name)
	if os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	} else {
		currentTime := time.Now()
		err = os.Chtimes(name, currentTime, currentTime)
		if err != nil {
			panic(err)
		}
	}
}

func FilterRead(items []Item) []Item {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	var filtered []Item
	for _, item := range items {
		_, err = os.Stat(homeDir + "/.feedline/read/" + sanitize.BaseName(item.URL))
		if os.IsNotExist(err) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
