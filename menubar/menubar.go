package menubar

import (
	"fmt"
	"gopkg.in/qml.v1"
)

func Initialize(win *qml.Window) {
	about(win)
}

func about(win *qml.Window) {
	win.ObjectByName("menu:help:about").On("triggered", func() {
		fmt.Println("Some useful stuff...")
	})
}