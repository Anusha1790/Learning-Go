package config

import (
	"html/template"
	"log"
)

// site-wide app configuration
// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
