package main

import (
	"./app"
	"./filemanager"
	"fmt"
	"gopkg.in/qml.v1"
)

func main() {
	err := qml.Run(initialize)
	if err != nil {
		fmt.Println(err)
	}
}

func initialize() error {
	context := app.NewContext()

	go context.ActionManager()

	context.Actions <- app.Action{
		Kind: filemanager.FILE_OPEN,
		Payload: nil,
	}

	context.Actions <- app.Action{
		Kind: filemanager.FILE_OPEN,
		Payload: nil,
	}

	return <- context.Exit
}
