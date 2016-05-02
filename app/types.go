package app

import (
	"../filemanager"
	"gopkg.in/qml.v1"
)

type Context struct {
	Engine qml.Engine
	Actions chan Action
	Exit chan error
	Files filemanager.Map
}

type Action struct {
	Kind int
	Payload interface{}
}
