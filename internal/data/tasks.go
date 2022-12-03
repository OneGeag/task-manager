package data

import "time"

type Task struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"-"`
	Expires time.Time `json:"-"`
}
