package menubar

import (
	"gopkg.in/qml.v1"
)

func Initialize(win *qml.Window, engine qml.Engine, filesChan chan string) {
	file(win, engine, filesChan)
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

func file(win *qml.Window, engine qml.Engine, filesChan chan string) {
	fileDialog := win.ObjectByName("fileDialog")

	fileDialog.On("accepted", func() {
		filesChan <- fileDialog.String("fileUrl")
	})

	win.ObjectByName("menu:file:open").On("triggered", func() {
		fileDialog.Call("open")
	})
}