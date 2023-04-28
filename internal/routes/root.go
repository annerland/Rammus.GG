package routes

import (
	"io"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	n, err := io.WriteString(w, "OK")

	isErr := logErr(n, err)

	if isErr {
		return
	}
}