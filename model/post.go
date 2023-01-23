package model

type Post struct {
	Id          string     `json:"id"`
	Text        string     `json:"text"`
	Image       string     `json:"image"`
	Likes       int        `json:"likes"`
	Tags        []string   `json:"tags"`
	PublishDate string     `json:"publishDate"`
	Link        string     `json:"link"`
	Comments    []*Comment `json:"comments,omitempty"`
}
