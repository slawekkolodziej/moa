package editor

import (
	"bytes"
	"github.com/russross/blackfriday"
	"github.com/limetext/qml-go"
	"regexp"
	"text/template"
)

type Compiler func(qml.Object, qml.Object)

type Document struct {
	Name string
	Body string
}

func Initialize(win *qml.Window, tpl string, content []byte) {
	source := win.ObjectByName("source")
	target := win.ObjectByName("target")
	compile := getCompiler(tpl)

	source.Set("text", string(content))

	compile(source, target)
	watch(source, target, compile)
}

func watch(source qml.Object, target qml.Object, compile Compiler) {
	source.On("textChanged", func() {
		compile(source, target)
    })
}

func getCompiler(tpl string) Compiler {
	doc := template.Must(template.New("htmlDocument").Parse(tpl))

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
