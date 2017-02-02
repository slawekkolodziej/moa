package main

import (
	"./app"
	"./filemanager"
	"github.com/limetext/qml-go"
)

func main() {
	err := qml.Run(initialize)
	if err != nil {
		panic(err)
	}
}

func initialize() error {
	context, err := app.NewContext()
	if err != nil {
		panic(err)
	}

	go context.ActionManager()

	context.Actions <- app.Action{
		Kind: filemanager.FILE_OPEN,
		Payload: nil,
	}

	return <- context.Exit
}
