package routes

import "github.com/go-chi/chi/v5"

func Register(r *chi.Mux) {
	// here I'm gonna call all the routes

	r.Get("/", rootHandler)

	r.NotFound(notFoundHandler)

	r.MethodNotAllowed(notAllowedHandler)
}
