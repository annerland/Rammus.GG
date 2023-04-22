package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))

		if err != nil {
			fmt.Printf("Error write")

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)

		_, err := w.Write([]byte("Route does not exist"))

		if err != nil {
			fmt.Printf("Error write")

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)

		_, err := w.Write([]byte("Method is not valid"))

		if err != nil {
			fmt.Printf("Error write")

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	})

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		fmt.Printf("Error starting the server")

		return
	}
}
