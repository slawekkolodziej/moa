package menubar

import (
	"fmt"
	"gopkg.in/qml.v1"
)

func Initialize(win *qml.Window, engine qml.Engine, filesChan chan string) {
	fileOpen(win, engine, filesChan)
	fileSave(win, engine)
	about(win, engine)
}

func about(win *qml.Window, engine qml.Engine) {
	win.ObjectByName("menu:help:about").On("triggered", func() {
		aboutComponent, err := engine.LoadFile("components/about.qml")
		if err == nil {
			aboutWindow := aboutComponent.CreateWindow(nil)
			aboutWindow.Show()
		}
	})
}

func fileOpen(win *qml.Window, engine qml.Engine, filesChan chan string) {
	fileDialog := win.ObjectByName("fileDialog")

	fileDialog.On("accepted", func() {
		filesChan <- fileDialog.String("fileUrl")
	})

	win.ObjectByName("menu:file:open").On("triggered", func() {
		fileDialog.Call("open")
	})
}

func fileSave(win *qml.Window, engine qml.Engine) {
	win.ObjectByName("menu:file:save").On("triggered", func() {
		fmt.Println("foo")
	})
}