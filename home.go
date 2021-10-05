package main

import "github.com/shoriwe/static/pkg/engine"

func HandleHome(e1 *engine.Engine) error {
	return e1.HandlePath("/index.html",
		func(e *engine.Engine) ([]byte, error) {
			index, indexCompileError := renderMarkdown("markdown/index.md")
			if indexCompileError != nil {
				return nil, indexCompileError
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

			navigationCssURL, navigationCssGetError := e.AssetURL("css/page.css")
			if navigationCssGetError != nil {
				return nil, navigationCssGetError
			}
			return e.RenderTemplate("page.html",
				map[string]string{
					"TITLE":          "Home",
					"MARKDOWN_STYLE": markdownCssURL,
					"PRISM_STYLE":    prismCssURL,
					"NAV_STYLE":      navigationCssURL,
					"BODY":           string(index),
					"PRISM_JS":       prismScriptURL,
					"MORE":           "",
				},
			)
		},
	)
}
