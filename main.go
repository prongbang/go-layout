package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var templates map[string]*template.Template

// Load templates on program initialisation
func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	pwd, _ := os.Getwd()
	templatesDir := pwd

	fmt.Println(templatesDir + "/public/views/*.html")

	layouts, err := filepath.Glob(templatesDir + "/public/views/*.html")
	if err != nil {
		log.Fatal(err)
	}

	// includes, err := filepath.Glob(templatesDir + "includes/*.tmpl")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		// files := append(includes, layout)
		// templates[filepath.Base(layout)] = template.Must(template.ParseFiles(files...))
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(layout))
	}

}

// renderTemplate is a wrapper around template.ExecuteTemplate.
func renderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {

	// hoge
	http.HandleFunc("/", hogeHandler)
	http.HandleFunc("/hoge", hogeHandler)

	// piyo
	// http.HandleFunc("/piyo", piyoHandler)

	http.ListenAndServe(":8080", nil)
}

// var hogeTmpl = template.Must(template.New("hoge").ParseFiles("base.html", "hoge.html"))

func hogeHandler(w http.ResponseWriter, r *http.Request) {
	// hogeTmpl.ExecuteTemplate(w, "base", "Hoge")
	renderTemplate(w, "index", nil)
}

// var piyoTmpl = template.Must(template.New("piyo").ParseFiles("base.html", "piyo.html"))

// func piyoHandler(w http.ResponseWriter, r *http.Request) {
// piyoTmpl.ExecuteTemplate(w, "base", "Piyo")
// }
