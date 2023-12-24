package stlos

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	stlbasic "github.com/kkkunny/stl/basic"
)

// TempDir 缓存目录地址
func TempDir() FilePath {
	return FilePath(os.TempDir())
}

const (
	// 随机名字长度
	randomNameLength = 10
)

// RandomTempFilePath 随机一个缓存文件地址
func RandomTempFilePath(prefix string) (FilePath, error) {
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		var nameBuffer strings.Builder
		for i := 0; i < randomNameLength; i++ {
			num := rune(rander.Intn(52))
			char := stlbasic.Ternary(num < 26, 'a'+num, 'A'+num-26)
			nameBuffer.WriteRune(char)
		}
		var name string
		if prefix != "" {
			name = fmt.Sprintf("%s_%s.tmp", prefix, nameBuffer.String())
		} else {
			name = fmt.Sprintf("%s.tmp", nameBuffer.String())
		}
		path := TempDir().Join(name)
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
func CreateTempFile(prefix string) (FilePath, *os.File, error) {
	path, err := RandomTempFilePath(prefix)
	if err != nil {
		return "", nil, err
	}
	file, err := os.Create(path.String())
	if err != nil {
		return "", nil, err
	}
	return path, file, nil
}

type TempFile struct {
	path FilePath
	*os.File
}

func (f *TempFile) Close() error {
	if err := f.File.Close(); err != nil {
		return err
	}
	return os.Remove(f.path.String())
}

// CreateTempFileWithCloser 创建一个带自动删除的缓存文件
func CreateTempFileWithCloser(prefix string) (*TempFile, error) {
	path, file, err := CreateTempFile(prefix)
	if err != nil {
		return nil, err
	}
	return &TempFile{
		path: path,
		File: file,
	}, nil
}
