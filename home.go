package main

import "github.com/shoriwe/static/pkg/engine"

func HandleHome(e1 *engine.Engine) error {
	return e1.HandlePath("/index.html",
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
}
