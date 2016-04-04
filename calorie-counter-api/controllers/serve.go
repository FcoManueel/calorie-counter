package controllers

import (
	"encoding/json"
	"github.com/FcoManueel/calorie-counter/calorie-counter-api/errors"
	"golang.org/x/net/context"
	"net/http"
)

// ParseBody message into a struct
func ParseBody(ctx context.Context, message interface{}, r *http.Request) error {
	if r.ContentLength == 0 {
		return errors.New(errors.BAD_REQUEST, "Empty request body")
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&message); err != nil {
		return errors.New(errors.INTERNAL_SERVER_ERROR, "Error while unmarshalling request body")
	}
	return nil
}

// ServeJSON uses a json encoder to serialize to the stream
func ServeJSON(ctx context.Context, w http.ResponseWriter, message interface{}) error {
	return json.NewEncoder(w).Encode(message)
}

// ServeError renders a JSON appError nicely or a 500 error for basic errors
func ServeError(ctx context.Context, w http.ResponseWriter, err error) {
	if err, ok := err.(errors.AppError); ok {
		w.WriteHeader(err.HTTPCode())
		ServeJSON(ctx, w, err)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
