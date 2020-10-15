package model

//Article struct
type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    int64 `json:"date"`
}
