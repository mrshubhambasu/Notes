package models

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Created string `json:"created"`
}
