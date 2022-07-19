package templates

import (
	"bytes"
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
	buffer := bytes.NewBuffer(nil)
	executeError := pageTemplate.Execute(buffer, p)
	if executeError != nil {
		return executeError
	}
	newBody := bytes.ReplaceAll(buffer.Bytes(), []byte(`href="/`), []byte(`href="/plasma/`))
	newBody = bytes.ReplaceAll(newBody, []byte(`src="/`), []byte(`src="/plasma/`))
	_, writeError := w.Write(newBody)
	return writeError
}
