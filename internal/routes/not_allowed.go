package routes

import (
	"fmt"
	"net/http"
)

func notAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)

	_, err := w.Write([]byte("Method is not valid"))

	if err != nil {
		fmt.Printf("Error write")

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}