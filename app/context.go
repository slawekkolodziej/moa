package app;

import (
    "../editor"
    "../filemanager"
    "../webengine"
    "fmt"
    "gopkg.in/qml.v1"
)

type Context struct {
    Engine qml.Engine
    Actions chan Action
    Exit chan error
    Files *filemanager.Map
    MenuBar *qml.Object
    Windows []*qml.Window
}

func NewContext() (Context, error) {
    context := Context{
        Engine: *qml.NewEngine(),
        Actions: make(chan Action, 1),
        Exit: make(chan error, 1),
        Files: filemanager.New(),
        Windows: make([]*qml.Window, 0),
    }

    menubar, err := context.NewMenubar();
    if err != nil {
        return context, err
    }
    context.MenuBar = menubar

    webengine.Initialize()
    go context.ActionManager()

    return context, nil
}

func (context Context) ActionManager() {
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
            win, err := context.NewWindow(file)
            if err != nil {
                panic(err)
            }
            context.Windows = append(context.Windows, win)

        case filemanager.FILE_SAVE:
            file := nextAction.Payload.(filemanager.File)
            fmt.Println("action type: FILE_SAVE", file)

        case filemanager.FILE_CLOSE:
            context.Files.Close(nextAction.Payload.(filemanager.File))

        case filemanager.FILE_OPEN_DIALOG:
            var win *qml.Window
            // fmt.Println(len(context.Windows))
            for i := 0; i < len(context.Windows); i++ {
                win = context.Windows[i]
                fmt.Println(win.Bool("active"))
            }
        }

        fmt.Println("total files opened: ", context.Files)

        if (context.Files.Total() == 0) {
            context.Exit <- nil
        }
    }
}

func (context Context) NewWindow(file filemanager.File) (*qml.Window, error) {
    appComponent, err := context.Engine.LoadFile("components/app.qml")
    if err != nil {
        return nil, err
    }

    content, err := file.Content();
    if err != nil {
        return nil, err
    }

    win := appComponent.CreateWindow(nil)
    editor.Initialize(win, htmlDocument, content)

    win.Set("title", file.Name)
    win.On("closing", func() {
        context.Actions <- Action{
            Payload: file,
            Kind: filemanager.FILE_CLOSE,
        }
    })
    win.Show()

    return win, nil
}