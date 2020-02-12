package logger

import (
	"github.com/pkg/errors"
	"os"
)

type FileWrite struct {
	file *os.File
}

// Write
func (fw *FileWrite) Write(data []byte) (n int, err error) {
	if fw == nil {
		return 0, errors.New("FileWrite is nil")
	}
	if fw.file == nil {
		return 0, errors.New("FileWrite file not opened")
	}
	n, e := fw.file.Write(data)
	return n, e
}
