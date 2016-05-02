package main

import (
	"./editor"
	"./filemanager"
	"./menubar"
	"./types"
	"./webengine"
	"fmt"
	"gopkg.in/qml.v1"
	// "os"
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
	err := qml.Run(app)
	if err != nil {
		fmt.Println(err)
	}
}

func app() error {
	context := newContext()

	context.Actions <- types.Action{
		Kind: filemanager.FILE_OPEN,
		Payload: nil,
	}

	return <- context.Exit
}

func actionManager(context types.AppContext) {
	for {
		nextAction := <- context.Actions

		switch nextAction.Kind {
		case filemanager.FILE_OPEN:
			var filePath *string = nil

			switch nextAction.Payload.(type) {
			case *string:
				filePath = nextAction.Payload.(*string)
			default:
				filePath = nil
			}

			file := context.Files.Open(filePath)
			fmt.Println("Open window: ", file.Name, file.Path)
			openWindow(context, file)

		case filemanager.FILE_SAVE:
			fmt.Println("action type: FILE_SAVE")
		case filemanager.FILE_CLOSE:
			fmt.Println("action type: FILE_CLOSE", nextAction.Payload, context.Files)
		}

		fmt.Println("total files opened: ", context.Files)

		if (context.Files.Total() == 0) {
			context.Exit <- nil
		}
	}
}

func openWindow(context types.AppContext, file filemanager.File) error {
	appComponent, err := context.Engine.LoadFile("components/app.qml")
	if err != nil {
		return err
	}

	content, err := file.Content();
	if err != nil {
		return err
	}

	fmt.Println("content: ", content)

	win := appComponent.CreateWindow(nil)
	menubar.Initialize(win, context);
	editor.Initialize(win, htmlDocument, content)

	win.Set("title", file.Name)
	win.On("closing", func() {
		// context.Actions <- types.Action{
			// Payload: filemanager.File{
			// 	Path: filePathPtr,
			// },
		// 	Kind: filemanager.FILE_CLOSE,
		// }
	})
	win.Show()

	return nil
}

func newContext() types.AppContext {
	var context types.AppContext;

	context.Engine = *qml.NewEngine()
	context.Actions = make(chan types.Action)
	context.Files = filemanager.New()
	context.Exit = make(chan error, 1)

	webengine.Initialize()

	go actionManager(context)

	return context
}