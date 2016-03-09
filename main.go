package main

import (
	"fmt"
	"gopkg.in/qml.v1"
	"github.com/russross/blackfriday"
	"text/template"
	"bytes"
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
	engine := qml.NewEngine()
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
	htmlDoc := template.Must(template.New("htmlDocument").Parse(`{{.Body}}`))

	return func(source qml.Object, target qml.Object) {
		var buf bytes.Buffer

		input := source.String("text")
		output := blackfriday.MarkdownBasic([]byte(input))

		htmlDoc.Execute(&buf, &Document{"test", string(output)})
		target.Call("loadHtml", buf.String())
	}
}