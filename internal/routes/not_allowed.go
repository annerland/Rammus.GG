package routes

import (
	"io"
	"net/http"
)

func notAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)

	n, err := io.WriteString(w, "Method is not valid.")

	isErr := logErr(n, err)

	if isErr {
		return
	}
}