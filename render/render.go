package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTemplate *template.Template
}

func Parse(filepath string) (Template, error) {
	htmlTemplate, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing went wrong %w", err)
	}
	return Template{
		htmlTemplate: htmlTemplate,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTemplate.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}
func RenderError(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}
