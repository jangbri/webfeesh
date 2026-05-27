package item

import "time"

type Item struct {
	ID          int64     `json:"id"`
	FeedID      int64     `json:"feed_id"`
	GUID        string    `json:"guid"`
	Link        string    `json:"link"`
	Title       string    `json:"title"`
	DateFetched time.Time `json:"date_fetched"`
	DateUpdated time.Time `json:"date_updated"`
}
