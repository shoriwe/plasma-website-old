package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/shoriwe/static/pkg/template"
	"io"
	"os"
	"strings"
)

func renderMarkdown(filePath string) ([]byte, error) {
	extensions := parser.CommonExtensions
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
	customRenderer := html.NewRenderer(
		html.RendererOptions{
			Flags: html.CommonFlags,
		},
	)
	contentAsHTML := markdown.ToHTML(contents, mdParser, customRenderer)
	newContentAsHTML := strings.ReplaceAll(string(contentAsHTML), "class=\"language-go\"", "class=\"language-go line-numbers\"")
	newContentAsHTML = strings.ReplaceAll(string(contentAsHTML), "class=\"language-ruby\"", "class=\"language-ruby line-numbers\"")
	t := template.NewTemplate(
		strings.NewReader(`
<article class="markdown-body">
	(( BODY ))
</article>`),
		map[string]string{
			"BODY": newContentAsHTML,
		},
	)
	return t.Compile()
}
