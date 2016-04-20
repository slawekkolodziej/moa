package main

import (
	"./editor"
	"./menubar"
	"./webengine"
	"fmt"
	"gopkg.in/qml.v1"
	"io/ioutil"
	"os"
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
	err := qml.Run(runApp)
	if err != nil {
		fmt.Println(err)
	}
}

func runApp() error {
	files := make(chan string)

	engine := qml.NewEngine()
	engine.On("quit", func() {
		os.Exit(0)
	})

	webengine.Initialize()

	go fileManager(*engine, files)

	err := openWindow(*engine, files, nil)
	if err != nil {
		return err
	}

	return nil
}

func openWindow(engine qml.Engine, files chan string, filePathPtr *string) error {
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
	win.Wait()

	return nil
}

func fileManager(engine qml.Engine, files chan string) {
	for {
		filePath := <- files
		err := openWindow(engine, files, &filePath)
		if err != nil {
			fmt.Println(err)
		}
	}
}