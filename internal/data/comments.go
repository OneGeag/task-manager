package data

import (
	"time"
)

type Comment struct {
	CommentID int64 `json:"commentid"`
	TaskID int64 `json:"taskid"`
	CreatedAt time.Time `json:"createdat"`
	Content string `json:"content"`
}
