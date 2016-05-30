package app

import (
	"../filemanager"
	"gopkg.in/qml.v1"
)

func (context Context) NewMenubar() (*qml.Object, error) {
	component, err := context.Engine.LoadFile("components/menubar.qml")
	if err != nil {
		return nil, err
	}

	menubar := component.Create(nil);

	// fileOpen(menubar, context)
	// fileSave(menubar, context, file)
	about(menubar, context)

	return &menubar, nil
}

func about(obj qml.Object, context Context) {
	obj.ObjectByName("menu:help:about").On("triggered", func() {
		aboutComponent, err := context.Engine.LoadFile("components/about.qml")
		if err == nil {
			aboutWindow := aboutComponent.CreateWindow(nil)
			aboutWindow.Show()
		}
	})
}

func fileOpen(obj qml.Object, context Context) {
	fileDialog := obj.ObjectByName("fileDialog")

	fileDialog.On("accepted", func() {
		fileUrl := fileDialog.String("fileUrl")
		context.Actions <- Action{
			Kind: filemanager.FILE_OPEN,
			Payload: &fileUrl,
		}
	})

	obj.ObjectByName("menu:file:open").On("triggered", func() {
		fileDialog.Call("open")
	})
}

func fileSave(obj qml.Object, context Context, file filemanager.File) {
	obj.ObjectByName("menu:file:save").On("triggered", func() {
		context.Actions <- Action{
			Kind: filemanager.FILE_SAVE,
			Payload: file,
		}
	})
}