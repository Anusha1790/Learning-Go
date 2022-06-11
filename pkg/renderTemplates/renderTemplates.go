package renderTemplates

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	container := template.Must(template.ParseFiles(
		"./templates/"+tmpl,
		"./templates/"+"base"+".layout.gohtml",
	))
	err := container.Execute(w, nil)
	if err != nil {
		log.Println("Error parsing templates:", err)
		return
	}
}
