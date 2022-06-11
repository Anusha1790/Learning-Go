package renderTemplates

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// FuncMap is a Map of Functions that I can use in a template
// create our own functions and pass those to our template
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err) // die here bcs if no templateCache, we don't have any pages to show so no use going furthut
	}

	parsedTemplate, ok := templateCache[tmpl]
	// if tmpl doesn't exist. Like instead of about.page.gohtml, if we receive yellow.page.gohtml which does not exist, then die
	if !ok {
		log.Fatal()
	}

	/* (old code: for reference) parse, then execute
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/"+"base"+".layout.gohtml")
	err = parsedTemplate.Execute(w, nil)
	*/

	// parse, then execute the parsed files
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error parsing templates:", err)
		return
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// tell the cache to get all files at a particular location
	myCache := map[string]*template.Template{}

	// get all the pages:
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	// loop through the *page.gohtmls
	for _, page := range pages {
		// get the actual base name of the page, extract the actual base name from the full path (page)
		name := filepath.Base(page)

		fmt.Println("page is currently", page)
		// create a Template Set based upon "page", passed it in empty variable functions
		// that will eventually have some functions in there which we are going to use
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Does this template match any particular layout?
		// now lets look for the existence of layouts in this particular folder called templates,
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		// if any match is found, then length of matches will be greater than 0:
		if len(matches) > 0 {
			// parse the layout:
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		// add the ts to the cache
		// eg. for about.page.gohtml, the map will give a fully parsed and ready to use template, ts
		myCache[name] = ts

		// ts is already parsed. Ab bas parsed to .Execute karna padhta he
	}

	return myCache, nil
}
