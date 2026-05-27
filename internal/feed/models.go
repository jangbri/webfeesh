package feed

type Feed struct {
	ID     int64  `json:"id"`
	ListID int64  `json:"list_id"`
	Title  string `json:"title"`
	Link   string `json:"link"`
}
