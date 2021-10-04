package main

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/shoriwe/static/pkg/engine"
	"github.com/shoriwe/static/pkg/template"
	"io"
	"os"
	"strings"
)

func HandleError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "\u001b[31m%s\u001b[0m\n", err)
	os.Exit(2)
}

func PrintUsage() {
	_, _ = fmt.Fprintf(os.Stderr, "%s OUTPUT_DIRECTORY\n", os.Args[0])
	os.Exit(1)
}

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
	scripts, loadingError = engine.LoadScripts("scripts", "js")
	if loadingError != nil {
		HandleError(loadingError)
	}
	assets, loadingError = engine.LoadAssets("assets", "static")
	if loadingError != nil {
		HandleError(loadingError)
	}
	www := engine.NewEngine(templates, scripts, assets)

	pathHandleError := www.HandlePath("/index.html",
		func(e *engine.Engine) ([]byte, error) {
			index, indexCompileError := renderMarkdown(e, "markdown/index.md")
			if indexCompileError != nil {
				return nil, indexCompileError
			}
			navigationCssURL, navigationCssGetError := e.AssetURL("css/page.css")
			if navigationCssGetError != nil {
				return nil, navigationCssGetError
			}
			return e.RenderTemplate("page.html",
				map[string]string{
					"NAV_STYLE": navigationCssURL,
					"BODY":      string(index),
				},
			)
		},
	)
	if pathHandleError != nil {
		HandleError(pathHandleError)
	}

	pathHandleError = www.HandlePath("/playground.html",
		func(e *engine.Engine) ([]byte, error) {
			playgroundCssURL, playgroundCssGetError := e.AssetURL("css/playground.css")
			if playgroundCssGetError != nil {
				return nil, playgroundCssGetError
			}
			markdownCssURL, markdownCssGetError := e.AssetURL("vendor/css/github-markdown.css")
			if markdownCssGetError != nil {
				return nil, markdownCssGetError
			}
			prismCssURL, prismCssGetError := e.AssetURL("vendor/prism/prism.css")
			if prismCssGetError != nil {
				return nil, prismCssGetError
			}
			prismScriptURL, prismScriptGetError := e.AssetURL("vendor/prism/prism.js")
			if prismScriptGetError != nil {
				return nil, prismScriptGetError
			}
			prismLiveCssURL, prismLiveCssGetError := e.AssetURL("vendor/prism/prism-live.css")
			if prismLiveCssGetError != nil {
				return nil, prismLiveCssGetError
			}
			prismLiveScriptURL, prismLiveURlGetError := e.AssetURL("vendor/prism/prism-live.js")
			if prismLiveURlGetError != nil {
				return nil, prismLiveURlGetError
			}
			playground, playgroundCompileError := e.RenderTemplate("playground.html",
				map[string]string{
					"PLAYGROUND_STYLE":  playgroundCssURL,
					"MARKDOWN_STYLE":    markdownCssURL,
					"PRISM_STYLE":       prismCssURL,
					"PRISM_LIVE_STYLE":  prismLiveCssURL,
					"PRISM_SCRIPT":      prismScriptURL,
					"PRISM_LIVE_SCRIPT": prismLiveScriptURL,
				},
			)
			if playgroundCompileError != nil {
				return nil, playgroundCompileError
			}
			navigationCssURL, navigationCssGetError := e.AssetURL("css/page.css")
			if navigationCssGetError != nil {
				return nil, navigationCssGetError
			}
			return e.RenderTemplate("page.html",
				map[string]string{
					"NAV_STYLE": navigationCssURL,
					"BODY":      string(playground),
				},
			)
		},
	)
	if pathHandleError != nil {
		HandleError(pathHandleError)
	}

	generationError := www.Generate(os.Args[1])
	if generationError != nil {
		HandleError(generationError)
	}
}
