package home

import (
	_ "embed"
	"github.com/shoriwe/plasma-website/templates"
	"github.com/spf13/afero"
	"html/template"
)

//go:embed index.html
var script template.HTML

func Write(output afero.Fs) error {
	file, createError := output.Create("index.html")
	if createError != nil {
		return createError
	}
	defer file.Close()
	return templates.Render(
		templates.Page{
			Title: "Home",
			Body:  script,
		},
		file,
	)
}
