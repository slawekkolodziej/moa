package menubar

import (
	"gopkg.in/qml.v1"
)

func Initialize(win *qml.Window, engine qml.Engine) {
	file(win, engine)
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

func file(win *qml.Window, engine qml.Engine) {
	fileDialog := win.ObjectByName("fileDialog")

	win.ObjectByName("menu:file:open").On("triggered", func() {
		fileDialog.Call("open")
	})
}