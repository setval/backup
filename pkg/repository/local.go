package repository

import (
	"io"
	"os"
	"path/filepath"
)

type local struct {
	dir string
}

func NewLocal(dir string) (*local, error) {
	return &local{
		dir: dir,
	}, nil
}

func (l *local) Upload(path string, reader io.Reader) error {
	if err := os.MkdirAll(filepath.Join(l.dir, filepath.Dir(path)), os.ModePerm); err != nil {
		return err
	}
	out, err := os.Create(filepath.Join(l.dir, path))
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, reader)
	return err
}
