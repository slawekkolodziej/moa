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
	Engine qml.Engine
	Actions chan Action
	Exit chan error
	Files []*string
}

type Action struct {
	File *string
	Kind int
}