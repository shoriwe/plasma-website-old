package docs

import (
	_ "embed"
	"github.com/shoriwe/plasma-website/templates"
	"github.com/spf13/afero"
	"html/template"
)

var (
	//go:embed documentation.html
	documentation template.HTML
)

func Write(output afero.Fs) error {
	createDirectoryError := output.MkdirAll("docs", 0777)
	if createDirectoryError != nil {
		return createDirectoryError
	}
	file, creationError := output.Create("docs/index.html")
	if creationError != nil {
		return creationError
	}
	defer file.Close()
	return templates.Render(
		templates.Page{
			Title: "Playground",
			Body:  documentation,
		},
		file,
	)
}
