package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// new router
	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.addIPToContext)
	mux.Use(app.Session.LoadAndSave)

	// register routes
	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)
	mux.Get("/user/profile", app.Profile)

	// static assets
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
