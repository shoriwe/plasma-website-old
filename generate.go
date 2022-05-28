package main

import (
	"fmt"
	"github.com/shoriwe/static/pkg/engine"
	"os"
)

func HandleError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "\u001b[31m%s\u001b[0m\n", err)
	os.Exit(2)
}

func PrintUsage() {
	_, _ = fmt.Fprintf(os.Stderr, "%s OUTPUT_DIRECTORY\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		PrintUsage()
	}
	var (
		templates    engine.Templates
		scripts      engine.Scripts
		assets       engine.Assets
		loadingError error
	)
	templates, loadingError = engine.LoadTemplates("templates")
	if loadingError != nil {
		HandleError(loadingError)
	}
	scripts, loadingError = engine.LoadScripts("scripts", "plasma/js")
	if loadingError != nil {
		HandleError(loadingError)
	}
	assets, loadingError = engine.LoadAssets("assets", "plasma/static")
	if loadingError != nil {
		HandleError(loadingError)
	}
	www := engine.NewEngine(templates, scripts, assets)

	pathHandleError := HandleHome(www)
	if pathHandleError != nil {
		HandleError(pathHandleError)
	}
	pathHandleError = HandlePlayground(www)
	if pathHandleError != nil {
		HandleError(pathHandleError)
	}
	pathHandleError = HandleDocumentation(www)
	if pathHandleError != nil {
		HandleError(pathHandleError)
	}

	generationError := www.Generate(os.Args[1])
	if generationError != nil {
		HandleError(generationError)
	}
}
