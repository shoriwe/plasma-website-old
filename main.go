package main

import (
	"github.com/shoriwe/plasma-website/articles"
	"github.com/shoriwe/plasma-website/home"
	"github.com/shoriwe/plasma-website/playground"
	"github.com/shoriwe/plasma-website/static"
	"github.com/spf13/afero"
	"log"
	"os"
)

var (
	output afero.Fs
)

func init() {
	output = afero.NewBasePathFs(afero.NewOsFs(), "www/plasma")
}

func main() {
	_ = os.RemoveAll(os.Args[1])
	writeStaticError := static.Write(output)
	if writeStaticError != nil {
		log.Fatal(writeStaticError)
	}
	writeDocumentationError := articles.Write(output)
	if writeDocumentationError != nil {
		log.Fatal(writeDocumentationError)
	}
	writeHomeError := home.Write(output)
	if writeHomeError != nil {
		log.Fatal(writeHomeError)
	}
	writePlaygroundError := playground.Write(output)
	if writePlaygroundError != nil {
		log.Fatal(writePlaygroundError)
	}
}
