package vocabulary

import (
	"os"
	"path/filepath"
)

type FileHandler interface {
	Read(dir, filename string) (string, error)
	Write(dir, filename, contents string) error
}

type Filesystem struct{}

func (*Filesystem) Read(dir, filename string) (string, error) {
	path := filepath.Join(dir, filename)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (*Filesystem) Write(dir, filename, contents string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	path := filepath.Join(dir, filename)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(contents)
	return err
}
