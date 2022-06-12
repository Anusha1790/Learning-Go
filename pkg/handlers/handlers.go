package handlers

import (
	"net/http"

	"github.com/Anusha1790/go-course/pkg/config"
	"github.com/Anusha1790/go-course/pkg/renderTemplates"
)

// handlers pkg also should have the configuration information

// Repo is the Repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// in order for a function to respond to a request from a web browser,
// it has to handle two parameters: w, r
// Home is the home page handler
func (rep *Repository) Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates.RenderTemplate(w, "home.page.gohtml")
}

func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	renderTemplates.RenderTemplate(w, "about.page.gohtml")

}
