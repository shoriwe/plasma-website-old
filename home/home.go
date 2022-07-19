package home

import (
	_ "embed"
	"github.com/shoriwe/plasma-website/markdown"
	"github.com/shoriwe/plasma-website/templates"
	"github.com/spf13/afero"
	"io"
	"net/http"
)

const readmeURL = "https://raw.githubusercontent.com/shoriwe/gplasma/main/README.md"

func Write(output afero.Fs) error {
	file, createError := output.Create("index.html")
	if createError != nil {
		return createError
	}
	defer file.Close()
	response, requestError := http.Get(readmeURL)
	if requestError != nil {
		return requestError
	}
	defer response.Body.Close()
	readme, readError := io.ReadAll(response.Body)
	if readError != nil {
		return readError
	}
	return templates.Render(
		templates.Page{
			Title: "Home",
			Body:  markdown.Render(readme),
		},
		file)
}
