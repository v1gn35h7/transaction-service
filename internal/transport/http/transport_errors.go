package http

import (
	"context"
	"encoding/json"
	"net/http"
)

type jsonError struct {
	Message string              `json:"message"`
	Errors  []map[string]string `json:"errors,omitempty"`
}

func baseError(err string) []map[string]string {
	return []map[string]string{map[string]string{
		"name":   "base",
		"reason": err},
	}
}

// encode error and status header to the client
func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	type validationError interface {
		error
		Invalid() []map[string]string
	}
	if e, ok := err.(validationError); ok {
		ve := jsonError{
			Message: "Validation Failed",
			Errors:  e.Invalid(),
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		enc.Encode(ve)
		return
	}

	type notFoundError interface {
		error
		IsNotFound() bool
	}
	if e, ok := err.(notFoundError); ok {
		je := jsonError{
			Message: "Resource Not Found",
			Errors:  baseError(e.Error()),
		}
		w.WriteHeader(http.StatusNotFound)
		enc.Encode(je)
		return
	}

	type existsError interface {
		error
		IsExists() bool
	}
	if e, ok := err.(existsError); ok {
		je := jsonError{
			Message: "Resource Already Exists",
			Errors:  baseError(e.Error()),
		}
		w.WriteHeader(http.StatusConflict)
		enc.Encode(je)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	je := jsonError{
		Message: "Service Error",
		Errors:  baseError(err.Error()),
	}
	enc.Encode(je)
}
