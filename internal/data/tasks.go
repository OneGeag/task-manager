package data

import "time"

type Task struct {
	TaskID int64 `json:"taskId"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []string `json:"tags"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"-"`
	Expires time.Time `json:"-"`
	Version int32 `json:"version"`
}
