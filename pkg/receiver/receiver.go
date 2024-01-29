package receiver

import (
	"time"
)

func Refresh() {
	head := linkedChannels()
	for head != nil {
		head.refresh()
		head = head.next
	}
}

func Listen(interval time.Duration) {
	// Currently, this doesn't do anything to shut itself down cleanly,
	// however, I'm not sure it needs to. As long as Refresh() doesn't perform
	// any disk operations, I'm not sure there's anything to clean up.

	// TODO Update at a user-configured time interval
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			Refresh()
		}
	}
}
