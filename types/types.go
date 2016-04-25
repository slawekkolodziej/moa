package types

import (
	"gopkg.in/qml.v1"
)

const (
	FILE_OPEN = iota
	FILE_SAVE
	FILE_CLOSE
)

type AppContext struct {
	engine qml.Engine
	actions chan Action
	exit chan error
	files []*string
}

type Action struct {
	file *string
	kind int
}