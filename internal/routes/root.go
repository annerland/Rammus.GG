package routes

import (
	"io"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	n, err := io.WriteString(w, "OK")

	logErr(n, err)
}