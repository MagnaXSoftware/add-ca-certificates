package management

import (
	"io/fs"
	"os"
)

type fileSystem struct {}

func (f fileSystem) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func (f fileSystem) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (f fileSystem) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(name)
}

func (f fileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (f fileSystem) Sub(dir string) (fs.FS, error) {
	return os.DirFS(dir), nil
}


var _ fs.FS = fileSystem{}
var _ fs.ReadDirFS = fileSystem{}
var _ fs.ReadFileFS = fileSystem{}
var _ fs.SubFS = fileSystem{}
