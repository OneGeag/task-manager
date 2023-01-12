package data

import (
	"time"
)

type Comment struct {
	CommentID int64     `json:"commentId"`
	TaskID    int64     `json:"taskId"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
}
