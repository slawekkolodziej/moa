package main

import (
	"./editor"
	"./webengine"
	"fmt"
	"gopkg.in/qml.v1"
	"os"
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

func main() {
	err := qml.Run(runApp)
	if err != nil {
		fmt.Println(err)
	}
}

func runApp() error {
	engine := qml.NewEngine()
	engine.On("quit", func() {
		os.Exit(0)
	})

	webengine.Initialize()

	appComponent, err := engine.LoadFile("components/app.qml")
	if err != nil {
		return err
	}

	win := appComponent.CreateWindow(nil)

	editor.Initialize(win, htmlDocument)

	win.Show()
	win.Wait()

	return nil
}
