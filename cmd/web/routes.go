package main

import (
	"net/http"

	"github.com/Anusha1790/go-course/pkg/handlers"
)

func routes() http.Handler {

	// firstly, create a multiplexer, which is basically an http Handler
	mux := http.NewServeMux()

	// this function is LISTENING TO a request sent by a web browser
	mux.HandleFunc("/", handlers.Repo.Home)
	mux.HandleFunc("/about", handlers.Repo.About)

	return mux
}
