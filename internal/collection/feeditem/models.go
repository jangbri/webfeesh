package feeditem

import (
	"encoding/xml"
	"time"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string                `xml:"title"`
	Link        string                `xml:"link"`
	Description string                `xml:"description"`
	Items       []*CollectionFeedItem `xml:"item"`
}

type CollectionFeedItem struct {
	ID           int64
	CollectionID int64
	Title        int64     `xml:"title"`
	Link         string    `xml:"link"`
	Description  string    `xml:"description"`
	TimeCreated  time.Time `xml:"pubDate,omitempty"`
}
