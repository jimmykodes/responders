package responders

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"error"`
}

func Respond(w http.ResponseWriter, body any, code int) {
	if body != nil {
		if err := json.NewEncoder(w).Encode(body); err != nil {
			slog.Error("error writing body", slog.String("error", err.Error()))
		}
	}
	w.WriteHeader(code)
}

func OK(w http.ResponseWriter, body any) {
	Respond(w, body, http.StatusOK)
}

func Error(w http.ResponseWriter, msg string, code int) {
	Respond(w, ErrorResponse{Message: msg}, code)
}

func Created(w http.ResponseWriter, body any) {
	Respond(w, body, http.StatusCreated)
}

func NoContent(w http.ResponseWriter) {
	Respond(w, nil, http.StatusNoContent)
}
