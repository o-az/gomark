package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	toc "github.com/abhinav/goldmark-toc"
	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func readFile(filepath string) string {
	body, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(body)
}

func writeFile(filepath string, content string) {
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatalf("unable to wrote file: %v", err)
	}
	defer file.Close()
	_, err2 := file.WriteString("<html>" + content + "</html>")

	if err2 != nil {
		log.Fatalf("unable to wrote file: %v", err)
	}
	fmt.Println("done")
}

func main() {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
			html.WithXHTML(),
			html.WithHardWraps(),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithExtensions(
			emoji.Emoji,
			mathjax.MathJax,
			meta.Meta,
			extension.GFM,
			&toc.Extender{},
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),

				highlighting.WithFormatOptions(),
			),
		),
	)
	var buf bytes.Buffer
	if err := markdown.Convert([]byte(readFile("./test-file.md")), &buf); err != nil {
		panic(err)
	}
	writeFile("./test-file.html", buf.String())
	fmt.Print(buf.String())
}
