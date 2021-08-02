package management

import (
	"io/fs"
	"os"
	"path/filepath"
)

type fileSystem struct {
	prefix string
}

func (f *fileSystem) Open(name string) (fs.File, error) {
	return os.Open(filepath.Join(f.prefix, name))
}

func (f *fileSystem) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(filepath.Join(f.prefix, name), flag, perm)
}

func (f *fileSystem) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(filepath.Join(f.prefix, name))
}

func (f *fileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(f.prefix, name))
}

func (f *fileSystem) Sub(dir string) (fs.FS, error) {
	return &fileSystem{prefix: filepath.Join(f.prefix, dir)}, nil
}

var _ fs.FS = &fileSystem{}
var _ fs.ReadDirFS = &fileSystem{}
var _ fs.ReadFileFS = &fileSystem{}
var _ fs.SubFS = &fileSystem{}
