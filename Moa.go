package main

import (
	"./editor"
	// "./menubar"
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

const (
	FILE_OPEN = iota
	FILE_SAVE
	FILE_CLOSE
)

type action struct {
	file *string
	kind int
}


func main() {
	err := qml.Run(app)
	if err != nil {
		fmt.Println(err)
	}
}

func app() error {
	engine := qml.NewEngine()
	// engine.On("quit", func() {
	// 	os.Exit(0)
	// })

	webengine.Initialize()

	exit := make(chan error, 1)
	actions := make(chan action)
	files := make([]*string, 0)

	// go fileManager(*engine, files)

	// err := openWindow(*engine, files, nil)
	// if err != nil {
	// 	return err
	// }

	go actionManager(*engine, actions, files, exit)

	actions <- action{file: nil, kind: FILE_OPEN}

	return <- exit
}

func actionManager(engine qml.Engine, actions chan action, files []*string, exit chan error) {
	for {
		nextAction := <- actions

		switch nextAction.kind {
		case FILE_OPEN:
			files = append(files, nextAction.file)
			fmt.Println("action type: FILE_OPEN", nextAction.file)
			openWindow(engine, nextAction.file)
		case FILE_SAVE:
			fmt.Println("action type: FILE_SAVE")
		case FILE_CLOSE:
			fmt.Println("action type: FILE_CLOSE")
		}

		fmt.Println("total files opened: ", len(files))

		if (len(files) == 0) {
			exit <- nil
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

func openWindow(engine qml.Engine, filePathPtr *string) error {
	var fileName string
	var filePath string
	var fileContent []byte

	appComponent, err := engine.LoadFile("components/app.qml")

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

	menubar.Initialize(win, engine, files);
	editor.Initialize(win, htmlDocument, fileContent)

	win.Show()

	return nil
}