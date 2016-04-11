package main

import (
	"./webengine"
	"bytes"
	"fmt"
	"github.com/russross/blackfriday"
	"gopkg.in/qml.v1"
	"os"
	"regexp"
	"text/template"
)

const htmlDocument = `
	<!doctype html>
	<html>
	<head>
		<title>{{.Name}}</title>
	</head>
	<body>
		{{.Body}}
	</body>
	</html>
`

type Compiler func(qml.Object, qml.Object)

type Document struct {
	Name string
	Body string
}

func main() {
	err := qml.Run(runApp)
	if err != nil {
		fmt.Println(err)
	}
}

func runApp() error {
	webengine.Initialize()
	engine := qml.NewEngine()
	engine.On("quit", func() {
		os.Exit(0)
	})

	winComponent, err := engine.LoadFile("components/base.qml")
	if err != nil {
		return err
	}

	win := winComponent.CreateWindow(nil)
	source := win.ObjectByName("source")
	target := win.ObjectByName("target")

	compile := runCompiler()

	compile(source, target)
	watch(source, target, compile)

	win.Show()
	win.Wait()


	return nil
}

func watch(source qml.Object, target qml.Object, compile Compiler) {
	source.On("textChanged", func() {
		compile(source, target)
    })
}

func runCompiler() Compiler {
	doc := template.Must(template.New("htmlDocument").Parse(htmlDocument))

	return func(source qml.Object, target qml.Object) {
		// Read value from the QML input
		input := source.String("text")

		// Cast string to []byte, and then pass it through blackfriday markdown parser
		output := blackfriday.MarkdownCommon([]byte(input))

		// Convert formatted text into a proper document
		html := processHtml(string(output), doc);

		// Fill the target container with the document
		target.Call("loadHtml", html)
	}
}

func processHtml(markdown string, doc *template.Template) string {
	// Prepare bytes buffer for later use
	var buf bytes.Buffer

	// Regular expression for new line replacement
	newLineRe := regexp.MustCompile("\n+")

	// Do some simple replacement in the doc
	markdownStr := newLineRe.ReplaceAllString(markdown, "<br>")

	// Fill the 'doc' template with values store
	doc.Execute(&buf, &Document{"test", markdownStr})

	// Return formatted HTML
	return buf.String()
}