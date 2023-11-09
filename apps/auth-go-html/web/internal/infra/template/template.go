package template

import (
	"html/template"
	"net/http"
)

type Templater interface {
	Handler(fn func(http.ResponseWriter, *http.Request, Templater)) http.HandlerFunc
	Render(w http.ResponseWriter, tmpl string, data any)
	RenderErrors(w http.ResponseWriter, tmpl string, errors map[string]string)
	Template() *template.Template
}

func NewTemplater(pattern string) Templater {
	return &templater{
		templates: template.Must(template.ParseGlob(pattern)),
	}
}

type templater struct {
	templates *template.Template
}
type ErrorData struct {
	Errors map[string]string
}

func (t *templater) Handler(fn func(http.ResponseWriter, *http.Request, Templater)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, t)
	}
}

func (t *templater) Render(w http.ResponseWriter, tmpl string, data any) {
	err := t.templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RenderErrors implements Templater.
func (t *templater) RenderErrors(w http.ResponseWriter, tmpl string, errors map[string]string) {
	t.Render(w, tmpl, &ErrorData{Errors: errors})
}

func (t *templater) Template() *template.Template {
	return t.templates
}
