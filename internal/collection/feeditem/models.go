package feeditem

import "time"

type CollectionFeedItem struct {
	ID           int64
	CollectionID int64
	Title        int64     `xml:"title"`
	Link         string    `xml:"link"`
	Description  string    `xml:"description"`
	TimeCreated  time.Time `xml:"pubDate"`
}
