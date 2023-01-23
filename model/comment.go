package model

type Comment struct {
	Id          string `json:"id"`
	Message     string `json:"message"`
	Post        string `json:"post"`
	PublishDate string `json:"publishDate"`
}
