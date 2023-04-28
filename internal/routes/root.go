package routes

import (
	"net/http"
)



func rootHandler(w http.ResponseWriter, r *http.Request) {
	n, err := w.Write([]byte("OK"))

	isErr := logErr(n, err)

	if isErr {
		return
	}
}