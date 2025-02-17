package models

type Task struct {
	ID      int    `json:"id"`
	Header  string `json:"header"`
	Content string `json:"content"`
	IsDone  bool   `json:"is_done"`
}
