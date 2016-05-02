package filemanager

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
	files map[uint32]File
}

func New() Map {
	var m Map
	return m
}

func NewFile(path *string) File {
	return File{
		Path: path,
	}
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
	delete(m.files, file.Id)
}

func (m Map) NextId() uint32 {
	m.lastId += 1
	return m.lastId
}

func (m Map) Total() int {
	return len(m.files)
}