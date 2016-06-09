package app

import (
    "../filemanager"
)

type Action struct {
	Kind int
	Payload interface{}
}

func (action Action) Open(context *Context) {
	var filePath *string = nil

    switch action.Payload.(type) {
    case *string:
        filePath = action.Payload.(*string)
    default:
        filePath = nil
    }

    file := context.Files.Open(filePath)
    _, err := context.NewWindow(file)
    assert(err)
}

func (action Action) Close(context *Context) {
	context.Files.Close(action.Payload.(filemanager.File))
}

func (action Action) Save(context *Context) {
    file := action.Payload.(*filemanager.File)
    win := file.Window
    markdown := win.ObjectByName("source").String("text")

    if file.Path == nil {
        fileDialog := win.ObjectByName("saveFile")
        fileDialog.On("accepted", func() {
            fileUrl := fileDialog.String("fileUrl")
            file.SetPath(&fileUrl)
            assert(file.Save(markdown))
        })
        fileDialog.Call("open")
    } else {
        assert(file.Save(markdown))
    }
}

func assert(err error) {
    if err != nil {
        panic (err)
    }
}