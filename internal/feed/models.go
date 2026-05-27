package feed

type Feed struct {
	ID           int64  `json:"id"`
	CollectionID int64  `json:"collection_id"`
	Title        string `json:"title"`
	Link         string `json:"link"`
}
