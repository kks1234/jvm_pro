package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath: abs}
}

func (e *ZipEntry) readClass(clasName string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(e.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()

	for _, f := range reader.File {
		if f.Name == clasName {
			open, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			all, err := io.ReadAll(open)
			if err != nil {
				open.Close()
				return nil, nil, err
			}
			open.Close()
			return all, e, nil
		}
	}
	return nil, nil, errors.New("class not found : " + clasName)
}

func (e *ZipEntry) String() string {
	return e.absPath
}
