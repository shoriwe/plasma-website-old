package main

import (
	"fmt"
	"github.com/shoriwe/static/pkg/engine"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func renderDocumentation(documentationFileName string, e *engine.Engine) ([]byte, error) {
	documentationCssURL, documentationCssGetError := e.AssetURL("css/documentation.css")
	if documentationCssGetError != nil {
		return nil, documentationCssGetError
	}
	tableOfContents, tableOfContentsRenderError := renderMarkdown("markdown/documentation/table-of-contents.md")
	if tableOfContentsRenderError != nil {
		return nil, tableOfContentsRenderError
	}
	documentation, documentationRenderError := renderMarkdown(filepath.Join("markdown/documentation/articles", documentationFileName))
	if documentationRenderError != nil {
		return nil, documentationRenderError
	}
	return e.RenderTemplate("documentation.html",
		map[string]string{
			"DOCUMENTATION_STYLE": documentationCssURL,
			"TABLE_OF_CONTENTS":   string(tableOfContents),
			"DOCUMENTATION":       string(documentation),
		},
	)
}

func createNewHandler(articleName string) engine.ContentGenerator {
	return func(e *engine.Engine) ([]byte, error) {
		body, bodyRenderError := renderDocumentation(articleName, e)
		if bodyRenderError != nil {
			return nil, bodyRenderError
		}
		navigationCssURL, navigationCssGetError := e.AssetURL("css/page.css")
		if navigationCssGetError != nil {
			return nil, navigationCssGetError
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

		page, pageRenderError := e.RenderTemplate("page.html",
			map[string]string{
				"TITLE":          "Documentation",
				"NAV_STYLE":      navigationCssURL,
				"MARKDOWN_STYLE": markdownCssURL,
				"PRISM_STYLE":    prismCssURL,
				"BODY":           string(body),
				"PRISM_JS":       prismScriptURL,
				"MORE":           "",
			},
		)
		if pageRenderError != nil {
			return nil, pageRenderError
		}
		return e.MinifyHTML(page)
	}
}
func HandleDocumentation(e1 *engine.Engine) error {
	articles, articlesReadError := os.ReadDir("markdown/documentation/articles")
	if articlesReadError != nil {
		return articlesReadError
	}
	for _, article := range articles {
		_, file := path.Split(article.Name())
		fileSplit := strings.Split(file, ".")
		toHandlePath := fmt.Sprintf("/plasma/documentation/%s.html", fileSplit[0])
		handlePathError := e1.HandlePath(toHandlePath, createNewHandler(article.Name()))
		if handlePathError != nil {
			return handlePathError
		}
	}
	return nil
}
