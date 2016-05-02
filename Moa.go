package main

import (
	"./app"
	"./editor"
	"./filemanager"
	"./menubar"
	"./webengine"
	"fmt"
	"gopkg.in/qml.v1"
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
	err := qml.Run(initialize)
	if err != nil {
		fmt.Println(err)
	}
}

func initialize() error {
	context := NewContext()

	go context.ActionManager()

	context.Actions <- app.Action{
		Kind: filemanager.FILE_OPEN,
		Payload: nil,
	}

	return <- context.Exit
}


// func NewContext() app.Context {
// 	var context app.Context;

// 	context.Engine = *qml.NewEngine()
// 	context.Actions = make(chan app.Action)
// 	context.Exit = make(chan error, 1)
// 	context.Files = filemanager.New()

// 	webengine.Initialize()

// 	go context.ActionManager()

// 	return context
// }

// func (context app.Context) ActionManager {
// 	for {
// 		nextAction := <- context.Actions

// 		switch nextAction.Kind {
// 		case filemanager.FILE_OPEN:
// 			var filePath *string = nil

// 			switch nextAction.Payload.(type) {
// 			case *string:
// 				filePath = nextAction.Payload.(*string)
// 			default:
// 				filePath = nil
// 			}

// 			file := context.Files.Open(filePath)
// 			context.NewWindow(context, file)

// 		case filemanager.FILE_SAVE:
// 			fmt.Println("action type: FILE_SAVE")

// 		case filemanager.FILE_CLOSE:
// 			fmt.Println("action type: FILE_CLOSE", nextAction.Payload, context.Files)
// 		}

// 		fmt.Println("total files opened: ", context.Files)

// 		if (context.Files.Total() == 0) {
// 			context.Exit <- nil
// 		}
// 	}
// }

// func (context app.Context) NewWindow(file filemanager.File) error {
// 	appComponent, err := context.Engine.LoadFile("components/app.qml")
// 	if err != nil {
// 		return err
// 	}

// 	content, err := file.Content();
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("content: ", content)

// 	win := appComponent.CreateWindow(nil)
// 	menubar.Initialize(win, context);
// 	editor.Initialize(win, htmlDocument, content)

// 	win.Set("title", file.Name)
// 	win.On("closing", func() {
// 		// context.Actions <- Action{
// 			// Payload: filemanager.File{
// 			// 	Path: filePathPtr,
// 			// },
// 		// 	Kind: filemanager.FILE_CLOSE,
// 		// }
// 	})
// 	win.Show()

// 	return nil
// }