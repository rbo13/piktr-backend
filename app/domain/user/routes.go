package user

import "github.com/go-chi/chi"

func Routes(r chi.Router, handler Handler) chi.Router {

	r.Post("/create", handler.Create)
	r.Get("/", handler.Get)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", handler.GetByID)
	})
	return r
}
