package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/onegeag/your-task-manager/internal/data"
)

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string        `json:"title"`
		Description string        `json:"description"`
		Tags        []string      `json:"tags"`
		Expires     int           `json:"expires"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Bad input", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readTaskIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	task := data.Task{
		TaskID:    id,
		CreatedAt: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"task": task}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server have got some issue", http.StatusInternalServerError)
	}
}
