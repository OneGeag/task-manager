package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Create a new task")
}

func (app *application) showTaskHandler(w http.ResponseWriter, r *http.Request) {

}
