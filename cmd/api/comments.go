package main

import (
	"net/http"
	"time"

	"github.com/onegeag/your-task-manager/internal/data"
)

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented yet(\n"))
}

func (app *application) showCommentHandler(w http.ResponseWriter, r *http.Request) {
	taskId, err := app.readTaskIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	commentId, err := app.readCommentIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	comment := data.Comment {
		TaskID: taskId,
		CommentID: commentId,
		CreatedAt: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"comment": comment}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server have got some issue", http.StatusInternalServerError)
	}
}
