package filemanager

import (
	"path"
	"io/ioutil"
	"sync/atomic"
)

const (
	FILE_OPEN = iota
	FILE_OPEN_DIALOG
	FILE_SAVE
	FILE_CLOSE
)

type File struct {
	Id uint32
	Path *string
	Name string
}

type Map struct {
	LastId uint32
	Files map[uint32]File
}

func New() *Map {
	return &Map{
		LastId: 0,
		Files: map[uint32]File{},
	}
}

func NewFile(filePath *string) File {
	return File{
		Path: filePath,
	}
}

func (m *Map) Open(filePath *string) File {
	file := File{
		Id: m.NextId(),
		Path: formatFilePath(filePath),
	}

	if (file.Path == nil) {
		file.Name = "Untitled"
	} else {
		file.Name = path.Base(*file.Path)
	}

	m.Files[file.Id] = file

	return file
}

func (m *Map) Close(file File) {
	delete(m.Files, file.Id)
}

func (m *Map) NextId() uint32 {
	return atomic.AddUint32(&m.LastId, 1)
}

func (m Map) Total() int {
	return len(m.Files)
}

func (file File) Content() ([]byte, error) {
	var content []byte
	var err error

	if file.Path == nil {
		content = []byte("")
	} else {
		content, err = ioutil.ReadFile(*file.Path)
		if err != nil {
			return nil, err
		}
	}

	return content, nil
}

func formatFilePath(filePath *string) *string {
	if (filePath != nil) {
		tmp := *filePath
		if tmp[:7] == "file://" {
			tmp := tmp[7:]
			filePath = &tmp
		}
	}

	return filePath
}