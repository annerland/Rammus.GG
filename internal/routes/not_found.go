package routes

import (
	"io"
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)

	n, err := io.WriteString(w, "Route doesn't exist.")

	isErr := logErr(n, err)

	if isErr {
		return
	}
}