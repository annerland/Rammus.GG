package routes

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))

	if err != nil {
		fmt.Printf("Error write")

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}