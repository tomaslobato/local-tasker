package models

type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Stamp int64  `json:"stamp"`
	Id    uint   `json:"id"`
}
