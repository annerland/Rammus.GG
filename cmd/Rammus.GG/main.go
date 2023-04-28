package main

import (
	"fmt"
	"net/http"

	"github.com/annerland/Rammus.GG/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	routes.Register(r)

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Printf("Error starting the server")

		return
	}
}
