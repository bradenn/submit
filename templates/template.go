package templates

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var tmpl *template.Template

// load
// Generate root template repository
func load() {
	dir, _ := os.Getwd()
	pattern := filepath.Join(dir+"/templates/", "*.gohtml")
	tmpl = template.Must(template.ParseGlob(pattern))

}

func Write(template string, writer http.ResponseWriter, data interface{}) {
	start := time.Now()
	load()
	fmt.Printf("Template discovery: %s\n", time.Since(start))
	err := tmpl.ExecuteTemplate(writer, template, data)
	if err != nil {
		fmt.Println(err)
	}
}
