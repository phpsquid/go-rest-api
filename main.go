package main

import (
	"log"
	"net/http"
	"phpsquid/myapi/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/user", user.Routes())
	})

	return router
}

func main() {
	router := Routes()

	log.Fatal(http.ListenAndServe(":8080", router))
}
