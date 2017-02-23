package app

import (
	"../filemanager"
	"github.com/limetext/qml-go"
)

func (context *Context) NewMenubar() (*qml.Object, error) {
	component, err := context.Engine.LoadString("", componentsMenubarQml)
	if err != nil {
		return nil, err
	}

	menubar := component.Create(nil);

	fileOpen(menubar, context)
	fileSave(menubar, context)
	about(menubar, context)

	return &menubar, nil
}

func about(obj qml.Object, context *Context) {
	obj.ObjectByName("menu:help:about").On("triggered", func() {
		aboutComponent, err := context.Engine.LoadString("", componentsAboutQml)
		if err == nil {
			aboutWindow := aboutComponent.CreateWindow(nil)
			aboutWindow.Show()
		}
	})
}

func fileOpen(obj qml.Object, context *Context) {
	obj.ObjectByName("menu:file:open").On("triggered", func() {
		fileDialog := context.Active.Window.ObjectByName("openFile")
		fileDialog.Call("open")
	})
}

func fileSave(obj qml.Object, context *Context) {
	obj.ObjectByName("menu:file:save").On("triggered", func() {
		context.Actions <- Action{
			Kind: filemanager.FILE_SAVE,
			Payload: context.Active,
		}
	})
}