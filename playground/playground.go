package playground

import (
	_ "embed"
	"github.com/shoriwe/plasma-website/templates"
	"github.com/spf13/afero"
	"html/template"
)

//go:embed playground.html
var playground template.HTML

func Write(output afero.Fs) error {
	createDirectoryError := output.MkdirAll("playground", 0777)
	if createDirectoryError != nil {
		return createDirectoryError
	}
	file, creationError := output.Create("playground/index.html")
	if creationError != nil {
		return creationError
	}
	defer file.Close()
	return templates.Render(
		templates.Page{
			Title: "Playground",
			Body:  playground,
		},
		file,
	)
}
