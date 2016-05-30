package app

import (
	"../filemanager"
	"gopkg.in/qml.v1"
)

func (context Context) NewMenubar(win *qml.Window, file filemanager.File) {
	fileOpen(win, context)
	fileSave(win, context, file)
	about(win, context)
}

func about(win *qml.Window, context Context) {
	win.ObjectByName("menu:help:about").On("triggered", func() {
		aboutComponent, err := context.Engine.LoadFile("components/about.qml")
		if err == nil {
			aboutWindow := aboutComponent.CreateWindow(nil)
			aboutWindow.Show()
		}
	})
}

func fileOpen(win *qml.Window, context Context) {
	fileDialog := win.ObjectByName("fileDialog")

	fileDialog.On("accepted", func() {
		fileUrl := fileDialog.String("fileUrl")
		context.Actions <- Action{
			Kind: filemanager.FILE_OPEN,
			Payload: &fileUrl,
		}
	})

	win.ObjectByName("menu:file:open").On("triggered", func() {
		fileDialog.Call("open")
	})
}

func fileSave(win *qml.Window, context Context, file filemanager.File) {
	win.ObjectByName("menu:file:save").On("triggered", func() {
		context.Actions <- Action{
			Kind: filemanager.FILE_SAVE,
			Payload: file,
		}
	})
}