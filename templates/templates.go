package templates

import (
	_ "embed"
	"html/template"
	"io"
)

var (
	//go:embed page.html
	page string
)

type Page struct {
	Title string
	Body  template.HTML
}

func Render(p Page, w io.Writer) error {
	pageTemplate, parseError := template.New(p.Title).Parse(page)
	if parseError != nil {
		return parseError
	}
	return pageTemplate.Execute(w, p)
}
