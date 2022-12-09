package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/OneGeag/task-manager/internal/data"
)

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Description string `json:"description"`
		Tags []string `json:"tags"`
		Expires int `json:"expires"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Bad input", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	task := data.Task{
		ID: id,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"task": task}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server have got some issue", http.StatusInternalServerError)
	}
}
