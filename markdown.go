package main

import (
	"bytes"

	toc "github.com/abhinav/goldmark-toc"
	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func compileMarkdown(content string) string {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			extension.TaskList,
			extension.Strikethrough,
			extension.Typographer,
			emoji.Emoji,
			mathjax.MathJax,
			meta.Meta,
			&toc.Extender{},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithAttribute(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(),
		),
	)
	var buf bytes.Buffer
	if err := markdown.Convert([]byte(content), &buf); err != nil {
		panic(err)
	}
	return `<html><head><link rel='preconnect' href='https://fonts.googleapis.com'/><linkrel='preconnect'href='https://fonts.gstatic.com'crossOrigin='annonymous'/><linkhref='https://fonts.googleapis.com/css2?family=IBM+Plex+Sans:wght@100;400;700&family=Inter:wght@200;400;900&family=JetBrains+Mono:wght@200;400;800&display=swap'rel='stylesheet'/><script src='https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.1/highlight.min.js'></script><linkhref='https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.1/styles/tokyo-night-dark.min.css' rel='stylesheet'type='text/css'/><script>hljs.highlightAll();</script><style>*{scroll-behavior:smooth}::-webkit-scrollbar{display:none}body,html{max-width:100%;max-height:100%;overflow:scroll;margin:0;padding:0;font-family:'IBM Plex Sans',sans-serif}body{max-width:800px;margin:50px auto 0;color:#f7f7f7;background-color:#0f0f0f}a{color:#6495ed;text-decoration:none}code{font-family:JetBrains Mono,monospace;line-height:1.4}</style></head><body>` + buf.String() + `</body></html>`
}
