package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/shoriwe/static/pkg/engine"
	"github.com/shoriwe/static/pkg/template"
	"io"
	"os"
	"strings"
)

func renderMarkdown(e *engine.Engine, filePath string) ([]byte, error) {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	mdParser := parser.NewWithExtensions(extensions)
	file, openError := os.Open(filePath)
	if openError != nil {
		return nil, openError
	}
	defer file.Close()
	contents, readError := io.ReadAll(file)
	if readError != nil {
		return nil, readError
	}
	contentAsHTML := markdown.ToHTML(contents, mdParser, nil)
	t := template.NewTemplate(
		strings.NewReader(`
<article class="markdown-body">
	(( BODY ))
</article>`),
		map[string]string{
			"BODY": string(contentAsHTML),
		},
	)
	return t.Compile()
}
