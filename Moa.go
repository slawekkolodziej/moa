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
	var context appContext;

	context.engine = *qml.NewEngine()
	context.actions = make(chan action)
	context.files = make([]*string, 0)
	context.exit = make(chan error, 1)
	// engine.On("quit", func() {
	// 	os.Exit(0)
	// })

	webengine.Initialize()

	// go fileManager(*engine, files)

	// err := openWindow(*engine, files, nil)
	// if err != nil {
	// 	return err
	// }

	go actionManager(context)

	context.actions <- action{file: nil, kind: FILE_OPEN}

	return <- context.exit
}

func actionManager(context appContext) {
	for {
		nextAction := <- context.actions

		switch nextAction.kind {
		case FILE_OPEN:
			context.files = append(context.files, nextAction.file)
			openWindow(context, nextAction.file)
		case FILE_SAVE:
			fmt.Println("action type: FILE_SAVE")
		case FILE_CLOSE:
			fmt.Println("action type: FILE_CLOSE")
		}

		fmt.Println("total files opened: ", len(context.files))

		if (len(context.files) == 0) {
			context.exit <- nil
		}
	}
}

// func fileManager(engine qml.Engine, files chan string) {
// 	for {
// 		filePath := <- files
// 		err := openWindow(engine, files, &filePath)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}
// }

func openWindow(context appContext, filePathPtr *string) error {
	var fileName string
	var filePath string
	var fileContent []byte

	appComponent, err := context.engine.LoadFile("components/app.qml")

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
	win.Set("title", fileName)

	menubar.Initialize(win, context);
	editor.Initialize(win, htmlDocument, fileContent)

	win.Show()

	return nil
}