package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anusha1790/go-course/pkg/config"
	"github.com/Anusha1790/go-course/pkg/handlers"
	"github.com/Anusha1790/go-course/pkg/renderTemplates"
)

const portNumber string = ":8080"

func main() {

	var app config.AppConfig

	templateCache, err := renderTemplates.CreateTemplateCache()
	if err != nil {
		log.Fatal(err) // die here bcs if no templateCache, we don't have any pages to show so no use going furthut
	}

	app.TemplateCache = templateCache

	// send app to renderTemplates using a function present in renderTemplates:
	renderTemplates.GetApp(app)

	// this function is LISTENING TO a request sent by a web browser
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_, _ = fmt.Printf("Starting aplication on port %s\n", portNumber)

	// LISTEN FOR requests from web browsers
	// start a web server that listens for requests in Go:
	log.Fatal(http.ListenAndServe(portNumber, nil)) // nil for Handler bcs http.HandleFunk is already doing the work for me

	// the first 1024 ports on any computer systems are privileged
}

// for making go.mod file: go mod init <path_name>

/*
when you move handlers to different file, go run main.go won't work, instead do:
go run *.go
*/
