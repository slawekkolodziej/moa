package app;

import (
    "../editor"
    "../filemanager"
    "../webengine"
    "github.com/limetext/qml-go"
)

type Context struct {
    active chan *filemanager.File
    Actions chan Action
    Active *filemanager.File
    Engine qml.Engine
    Exit chan error
    Files *filemanager.Map
    MenuBar *qml.Object
}

func NewContext() (Context, error) {
    context := Context{
        active: make(chan *filemanager.File, 1),
        Actions: make(chan Action, 1),
        Engine: *qml.NewEngine(),
        Exit: make(chan error, 1),
        Files: filemanager.New(),
    }

    menubar, err := context.NewMenubar();
    if err != nil {
        return context, err
    }
    context.MenuBar = menubar

    webengine.Initialize()
    go context.SetActive()
    go context.ActionManager()

    return context, nil
}

func (context *Context) SetActive() {
    for {
        context.Active = <- context.active
    }
}

func (context *Context) ActionManager() {
    for {
        nextAction := <- context.Actions

        switch nextAction.Kind {
        case filemanager.FILE_OPEN:
            nextAction.Open(context)

        case filemanager.FILE_SAVE:
            nextAction.Save(context)

        case filemanager.FILE_CLOSE:
            nextAction.Close(context)
        }

        if (context.Files.Total() == 0) {
            context.Exit <- nil
        }
    }
}

func (context *Context) NewWindow(file filemanager.File) (*qml.Window, error) {
    appComponent, err := context.Engine.LoadString("", componentsAppQml)
    if err != nil {
        return nil, err
    }

    content, err := file.Content();
    if err != nil {
        return nil, err
    }

    win := appComponent.CreateWindow(nil)
    file.Window = win
    editor.Initialize(win, htmlDocument, content)

    win.Set("title", file.Name)
    win.On("closing", func() {
        context.Actions <- Action{
            Payload: file,
            Kind: filemanager.FILE_CLOSE,
        }
    })
    win.On("activeChanged", func() {
        if win.Bool("active") == true {
            context.active <- &file
        }
    })
    fileDialog := win.ObjectByName("openFile")
    fileDialog.On("accepted", func() {
        fileUrl := fileDialog.String("fileUrl")
        context.Actions <- Action{
            Kind: filemanager.FILE_OPEN,
            Payload: &fileUrl,
        }
    })
    win.Show()

    if context.Active == nil {
        context.active <- &file
    }

    return win, nil
}