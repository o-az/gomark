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
	return `<!DOCTYPE html><link href=https://fonts.googleapis.com rel=preconnect><link href=https://fonts.gstatic.com rel=preconnect crossorigin><link href="https://fonts.googleapis.com/css2?family=IBM+Plex+Sans:wght@100;400;700&amp;family=Inter:wght@200;400;900&amp;family=JetBrains+Mono:wght@200;400;800&amp;display=swap"rel=stylesheet><script src=https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.1/highlight.min.js></script><link href=https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.1/styles/tokyo-night-dark.min.css rel=stylesheet><script>hljs.highlightAll()</script><style>body{max-width:800px;margin:50px auto 0;background-color:#0f0f0f}code{font-family:JetBrains Mono,monospace;line-height:1.4}</style>` + buf.String()
}
