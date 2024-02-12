package models

type Post struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
