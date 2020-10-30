package feed

import "time"

type Item struct {
	Subscription Subscription
	Label        string
	URL          string
	SanitizedURL string
	Date         time.Time
}
