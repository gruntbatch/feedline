package feed

import "time"

type Item struct {
	Label        string
	URL          string
	SanitizedURL string
	Date         time.Time
}
