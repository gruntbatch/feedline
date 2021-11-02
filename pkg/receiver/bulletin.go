package receiver

import (
	"errors"
	"os"
	"sort"
	"time"
)

type Bulletin struct {
	ItemInfo
	Date        time.Time
	ChannelInfo ItemInfo
}

func AllBulletins() []Bulletin {
	head := linkedChannels()

	var bulletins []Bulletin
	for head != nil {
		head.mu.RLock()
		bulletins = append(bulletins, head.bulletins...)
		head.mu.RUnlock()
		head = head.next
	}

	sort.Slice(bulletins, func(i, j int) bool {
		return bulletins[i].Date.After(bulletins[j].Date)
	})

	return bulletins
}

func UnreadBulletins() []Bulletin {
	return filterOutReadBulletins(AllBulletins())
}

func AllBulletinsFromSanitizedUrl(sanitizedURL string) ([]Bulletin, error) {
	head := linkedChannels()

	for head != nil {
		head.mu.RLock()
		defer head.mu.RUnlock()
		if head.Info.SanitizedURL == sanitizedURL {
			return head.bulletins, nil
		}
		head = head.next
	}

	return nil, errors.New("unable to find url")
}

func UnreadBulletinsFromSanitizedURL(sanitizedURL string) ([]Bulletin, error) {
	bulletins, err := AllBulletinsFromSanitizedUrl(sanitizedURL)
	if err != nil {
		return nil, err
	} else {
		return filterOutReadBulletins(bulletins), nil
	}
}

func filterOutReadBulletins(bulletins []Bulletin) []Bulletin {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	var filtered []Bulletin
	for _, bulletin := range bulletins {
		_, err = os.Stat(home + "/.feedline/read/" + bulletin.SanitizedURL)
		if os.IsNotExist(err) {
			filtered = append(filtered, bulletin)
		}
	}

	return filtered
}

func DismissBulletin(sanitizedURL string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	name := homeDir + "/.feedline/read/" + sanitizedURL

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
