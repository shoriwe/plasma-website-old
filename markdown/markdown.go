package markdown

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
)

const format = `<article class="markdown-body">%s</article>`

func Render(contents []byte) template.HTML {
	extensions := parser.CommonExtensions
	mdParser := parser.NewWithExtensions(extensions)
	customRenderer := html.NewRenderer(
		html.RendererOptions{
			Flags: html.CommonFlags,
		},
	)
	contentAsHTML := markdown.ToHTML(contents, mdParser, customRenderer)

	return template.HTML(fmt.Sprintf(format, string(contentAsHTML)))
}
