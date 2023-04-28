package routes

import (
	"fmt"
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)

	_, err := w.Write([]byte("Route does not exist"))

	if err != nil {
		fmt.Printf("Error write")

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}