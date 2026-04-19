package stlos

import (
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"

	stlbasic "github.com/kkkunny/stl/value"
)

const (
	// 随机名字长度
	randomNameLength = 10
)

// RandomTempFilePath 随机一个缓存文件地址
func RandomTempFilePath(prefix, ext string) (string, error) {
	for {
		var nameBuffer strings.Builder
		for i := 0; i < randomNameLength; i++ {
			num := rune(rand.IntN(52))
			char := stlbasic.If(num < 26, 'a'+num, 'A'+num-26)
			nameBuffer.WriteRune(char)
		}
		var name string
		if prefix != "" {
			name = fmt.Sprintf("%s_%s", prefix, nameBuffer.String())
		} else {
			name = nameBuffer.String()
		}
		if ext != "" {
			name += "." + ext
		}
		path := filepath.Join(os.TempDir(), name)
		exist, err := Exist(path)
		if err != nil {
			return "", err
		}
		if !exist {
			return path, nil
		}
	}
}

// CreateTempFile 创建一个缓存文件
func CreateTempFile(prefix, ext string) (string, *os.File, error) {
	path, err := RandomTempFilePath(prefix, ext)
	if err != nil {
		return "", nil, err
	}
	file, err := os.Create(path)
	if err != nil {
		return "", nil, err
	}
	return path, file, nil
}

type TempFile struct {
	path string
	*os.File
}

func (f *TempFile) Path() string {
	return f.path
}

func (f *TempFile) Close() error {
	if err := f.File.Close(); err != nil {
		return err
	}
	return os.Remove(f.path)
}

// CreateTempFileWithCloser 创建一个带自动删除的缓存文件
func CreateTempFileWithCloser(prefix, ext string) (*TempFile, error) {
	path, file, err := CreateTempFile(prefix, ext)
	if err != nil {
		return nil, err
	}
	return &TempFile{
		path: path,
		File: file,
	}, nil
}
