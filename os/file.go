package stlos

import (
	"io"
	"os"
)

// Exist 是否存在文件
func Exist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

// Copy 拷贝文件和目录
func Copy(from, to string) error {
	fromInfo, err := os.Stat(from)
	if err != nil {
		return err
	}

	if fromInfo.IsDir() {
		return os.CopyFS(to, os.DirFS(from))
	}

	fromFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	toFile, err := os.Create(to)
	if err != nil {
		return err
	}
	defer toFile.Close()

	_, err = io.Copy(toFile, fromFile)
	return err
}
