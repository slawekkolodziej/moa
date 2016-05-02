package filemanager

import (
	"path"
	"io/ioutil"
)

const (
	FILE_OPEN = iota
	FILE_SAVE
	FILE_CLOSE
)

type File struct {
	Id uint32
	Path *string
	Name string
}

type Map struct {
	lastId uint32
	files map[uint32]File
}

func New() Map {
	return Map{
		lastId: 0,
		files: make(map[uint32]File),
	}
}

func NewFile(filePath *string) File {
	return File{
		Path: filePath,
	}
}

func (m Map) Open(filePath *string) File {
	file := File{
		Id: m.NextId(),
		Path: formatFilePath(filePath),
	}

	if (file.Path == nil) {
		file.Name = "Untitled"
	} else {
		file.Name = path.Base(*file.Path)
	}

	m.files[file.Id] = file
	return file
}

func (m Map) Remove(file File) {
	delete(m.files, file.Id)
}

func (m Map) NextId() uint32 {
	m.lastId += 1
	return m.lastId
}

func (m Map) Total() int {
	return len(m.files)
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