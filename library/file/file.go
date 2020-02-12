package file

import (
	"fmt"
	"os"
)

// CreateFile
func CreateFile(pathStr string, fileName string) (*os.File, error) {
	var (
		err  error
		file *os.File
	)
	if pathStr, err = CreateDir(pathStr); nil != err {
		return file, err
	}
	return createFile(fmt.Sprintf("%s/%s", pathStr, fileName))
}

// OpenFile
func OpenFile(pathStr string, fileName string) (*os.File, error) {
	fileStr := fmt.Sprintf("%s/%s", pathStr, fileName)
	if !IsExist(fileStr) {
		return CreateFile(pathStr, fileName)
	}
	return os.OpenFile(fileStr, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0644)
}

// createFile
func createFile(fileStr string) (*os.File, error) {
	var (
		err  error
		file *os.File
	)
	if !IsExist(fileStr) {
		if file, err = os.Create(fileStr); nil != err {
			return file, err
		}
		if err = os.Chmod(fileStr, 0644); nil != err {
			return file, err
		}
	}
	return file, err
}

// CreateDir
func CreateDir(pathStr string) (string, error) {
	var err error
	if !IsExist(pathStr) {
		// 必须分成两步：先创建文件夹、再修改权限
		if err = os.MkdirAll(pathStr, 0755); nil != err {
			return pathStr, err
		}
		if err = os.Chmod(pathStr, 0755); nil != err {
			return pathStr, err
		}
	}
	return pathStr, err
}

// IsExist
func IsExist(filePathStr string) bool {
	if _, err := os.Stat(filePathStr); os.IsNotExist(err) {
		return false
	}
	return true
}
