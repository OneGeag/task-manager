package data

import "time"

type Task struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []string `json:"tags"`
	Status string `json:"status"`
	Progress Progress `json:"progress"`
	CreatedAt time.Time `json:"-"`
	Expires time.Time `json:"-"`
}
