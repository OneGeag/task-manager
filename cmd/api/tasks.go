package main

import (
	"fmt"
	"net/http"
)

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new task")
}

func (app *application) showTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the info of task %d\n", id)
}
