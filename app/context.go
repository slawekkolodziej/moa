package app;

import (
    "../editor"
    "../filemanager"
    "../webengine"
    "fmt"
    "gopkg.in/qml.v1"
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
            var filePath *string = nil

            switch nextAction.Payload.(type) {
            case *string:
                filePath = nextAction.Payload.(*string)
            default:
                filePath = nil
            }

            file := context.Files.Open(filePath)
            _, err := context.NewWindow(file)
            if err != nil {
                panic(err)
            }

        case filemanager.FILE_SAVE:
            file := nextAction.Payload.(*filemanager.File)
            markdown := file.Window.ObjectByName("source").String("text")
            err := file.Save(markdown)
            if err != nil {
                panic(err)
            }

        case filemanager.FILE_CLOSE:
            context.Files.Close(nextAction.Payload.(filemanager.File))
        }
        fmt.Println("total files opened: ", context.Files)

        if (context.Files.Total() == 0) {
            context.Exit <- nil
        }
    }
}

func (context *Context) NewWindow(file filemanager.File) (*qml.Window, error) {
    appComponent, err := context.Engine.LoadFile("components/app.qml")
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
        fmt.Println("window isActive:", win.Bool("active"))
        if win.Bool("active") == true {
            fmt.Println("setting is active...")
            context.active <- &file
        }
    })
    fileDialog := win.ObjectByName("fileDialog")
    fileDialog.On("accepted", func() {
        fileUrl := fileDialog.String("fileUrl")
        context.Actions <- Action{
            Kind: filemanager.FILE_OPEN,
            Payload: &fileUrl,
        }
    })
    win.Show()

    fmt.Println(context.Active)
    if context.Active == nil {
        fmt.Println("setting Active...")
        context.active <- &file
    }
    fmt.Println(context.Active)

    return win, nil
}