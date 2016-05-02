package types

import (
	"gopkg.in/qml.v1"
	"../filemanager"
)

type AppContext struct {
	Engine qml.Engine
	Actions chan Action
	Exit chan error
	Files filemanager.Map
}

type Action struct {
	Kind int
	Payload interface{}
}
