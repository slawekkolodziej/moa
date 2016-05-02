package filemanager

import (
	"fmt"
	"gopkg.in/qml.v1"
	"../types"
)

const (
	FILE_OPEN = iota
	FILE_SAVE
	FILE_CLOSE
)

type File struct {
	Id uint32
	Path *string
}

type Map struct {
	lastId uint32
	files map[uint32]File,
}

func New() Map {
	var m Map
	return m
}

func (m Map) Add(path string) uint32 {
	file := File{
		Id: m.NextId(),
		Path: &path,
	}

	m.files[file.Id] = file

	return file.Id
}

func (m Map) Remove(file File) {
	delete(m.files, file.id)
}

func (m Map) NextId() uint32 {
	m.lastId += 1
	return m.lastId
}