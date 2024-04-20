package models

type Video struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Content     string `json:"content"`
	UserID      int    `json:"user_id"`
}

type CreateVideoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	UserID      int    `json:"user_id"`
}
