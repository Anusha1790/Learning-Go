package handlers

import (
	"net/http"

	"github.com/Anusha1790/go-course/pkg/renderTemplates"
)

// in order for a function to respond to a request from a web browser,
// it has to handle two parameters: w, r
// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates.RenderTemplate(w, "home.page.gohtml")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates.RenderTemplate(w, "about.page.gohtml")

}
