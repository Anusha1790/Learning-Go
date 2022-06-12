package main

import (
	"net/http"

	"github.com/Anusha1790/go-course/pkg/handlers"
	"github.com/go-chi/chi"
)

func routes() http.Handler {

	mux := chi.NewRouter()
	//	r.Use(middleware.Logger)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// create fileServer: a place to go get static files from
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
