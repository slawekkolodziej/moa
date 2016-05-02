package menubar

import (
	"../filemanager"
	"../types/action"
	"../app"
	"fmt"
	"gopkg.in/qml.v1"
)

func Initialize(win *qml.Window, context app.Context) {
	fileOpen(win, context)
	fileSave(win, context.Engine)
	about(win, context.Engine)
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

func fileOpen(win *qml.Window, context app.Context) {
	fileDialog := win.ObjectByName("fileDialog")

	fileDialog.On("accepted", func() {
		fileUrl := fileDialog.String("fileUrl")
		context.Actions <- action.Action{
			Kind: filemanager.FILE_OPEN,
			Payload: &fileUrl,
		}
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