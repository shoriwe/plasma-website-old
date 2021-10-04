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

func renderMarkdown(engine *engine.Engine, filePath string) ([]byte, error) {
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
	markdownCssURL, markdownCssGetError := engine.AssetURL("vendor/css/github-markdown.css")
	if markdownCssGetError != nil {
		return nil, markdownCssGetError
	}
	prismCssURL, prismCssGetError := engine.AssetURL("vendor/prism/prism.css")
	if prismCssGetError != nil {
		return nil, prismCssGetError
	}
	prismScriptURL, prismScriptGetError := engine.AssetURL("vendor/prism/prism.js")
	if prismScriptGetError != nil {
		return nil, prismScriptGetError
	}
	t := template.NewTemplate(
		strings.NewReader(`
<link href="(( MARKDOWN_STYLE ))" rel="stylesheet">
<link href="(( PRISM_STYLE ))" rel="stylesheet">
<article class="markdown-body">
	(( BODY ))
</article>
<script src="(( PRISM_JS ))"></script>`),
		map[string]string{
			"MARKDOWN_STYLE": markdownCssURL,
			"PRISM_STYLE":    prismCssURL,
			"BODY":           string(contentAsHTML),
			"PRISM_JS":       prismScriptURL,
		},
	)
	return t.Compile()
}
