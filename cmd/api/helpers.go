package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	bytes = append(bytes, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)

	return nil
}

func (app *application) readIDParam(r *http.Request, param string) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName(param), 10, 64)
	if err != nil || id < 0 {
		return 0, errors.New("invalid id parametr")

	}
	return id, nil
}

func (app *application) readTaskIDParam(r *http.Request) (int64, error) {
	return app.readIDParam(r, "taskid")
}

func (app *application) readCommentIDParam(r *http.Request) (int64, error) {
	return app.readIDParam(r, "commentid")
}
