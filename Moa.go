package main

import (
	"./editor"
	"./menubar"
	"./types"
	"./webengine"
	"fmt"
	"gopkg.in/qml.v1"
	"io/ioutil"
	// "os"
	"path"
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
	var context types.AppContext;

	context.Engine = *qml.NewEngine()
	context.Actions = make(chan types.Action)
	context.Files = make([]*string, 0)
	context.Exit = make(chan error, 1)
	// engine.On("quit", func() {
	// 	os.Exit(0)
	// })

	webengine.Initialize()

	go actionManager(context)

	context.Actions <- types.Action{
		File: nil,
		Kind: types.FILE_OPEN,
	}

	return <- context.Exit
}

func actionManager(context types.AppContext) {
	for {
		nextAction := <- context.Actions

		switch nextAction.Kind {
		case types.FILE_OPEN:
			context.Files = append(context.Files, nextAction.File)
			openWindow(context, nextAction.File)
		case types.FILE_SAVE:
			fmt.Println("action type: FILE_SAVE")
		case types.FILE_CLOSE:
			fmt.Println("action type: FILE_CLOSE", nextAction, context.Files)
		}

		fmt.Println("total files opened: ", len(context.Files))

		if (len(context.Files) == 0) {
			context.Exit <- nil
		}
	}
}

func openWindow(context types.AppContext, filePathPtr *string) error {
	var fileName string
	var filePath string
	var fileContent []byte

	appComponent, err := context.Engine.LoadFile("components/app.qml")

	if err != nil {
		return err
	}

	if filePathPtr == nil {
		fileName = "Untitled"
		fileContent = []byte("")
	} else {
		filePath = *filePathPtr

		if filePath[:7] == "file://" {
			filePath = filePath[7:]
		}

		fileName = path.Base(filePath)
		fileContent, err = ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}
	}

	win := appComponent.CreateWindow(nil)

	menubar.Initialize(win, context);
	editor.Initialize(win, htmlDocument, fileContent)

	win.Set("title", fileName)
	win.On("closing", func() {
		context.Actions <- types.Action{
			File: filePathPtr,
			Kind: types.FILE_CLOSE,
		}
	})
	win.Show()

	return nil
}