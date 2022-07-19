package main

import (
	"fmt"
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
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s OUTPUT_DIRECTORY", os.Args[0])
		os.Exit(0)
	}
	output = afero.NewBasePathFs(afero.NewOsFs(), os.Args[1])
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
