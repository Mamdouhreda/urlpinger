package requests

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// Home handles the root route.
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

// renderTemplate parses and executes the named template from the frontend directory.
func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Resolve template relative to container workdir (/app) where frontend is copied
	templatePath := filepath.Join("frontend", tmpl)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
