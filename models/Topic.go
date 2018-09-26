package models

type Topic struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
