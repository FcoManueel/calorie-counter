package controllers

import (
	"encoding/json"
	"errors"
	"golang.org/x/net/context"
	"net/http"
)

// ParseBody message into a struct
func ParseBody(message interface{}, r *http.Request) error {
	if r.ContentLength == 0 {
		return errors.New("Empty request body")
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&message); err != nil {
		return errors.New("Error while unmarshalling request body")
	}
	return nil
}

// ServeJSON uses a json encoder to serialize to the stream
func ServeJSON(ctx context.Context, w http.ResponseWriter, message interface{}) error {
	return json.NewEncoder(w).Encode(message)
}

// ServeError renders a JSON appError nicely or a 500 error for basic errors
func ServeError(ctx context.Context, w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
