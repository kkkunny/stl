package stlos

import (
	"io/fs"
	"os"
	"path/filepath"

	stlbasic "github.com/kkkunny/stl/hash"
)

// FilePath 文件路径
type FilePath string

func NewFilePath(path string) FilePath {
	return FilePath(path)
}

// GetWorkDirectory 获取工作目录
func GetWorkDirectory() (FilePath, error) {
	fp, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return FilePath(fp), nil
}

func (self FilePath) Default() FilePath {
	return ""
}

func (self FilePath) String() string {
	return string(self)
}

func (self FilePath) Clone() FilePath {
	return self
}

func (self FilePath) Equal(dst FilePath) bool {
	return self == dst
}

func (self FilePath) Hash() uint64 {
	return stlbasic.Hash(string(self))
}

func (self FilePath) Abs() (FilePath, error) {
	fp, err := filepath.Abs(string(self))
	if err != nil {
		return "", err
	}
	return FilePath(fp), nil
}

func (self FilePath) Base() string {
	return filepath.Base(string(self))
}

func (self FilePath) Clean() FilePath {
	return FilePath(filepath.Clean(string(self)))
}

func (self FilePath) Dir() FilePath {
	return FilePath(filepath.Dir(string(self)))
}

func (self FilePath) Ext() string {
	return filepath.Ext(string(self))
}

func (self FilePath) EvalSymlinks() (FilePath, error) {
	fp, err := filepath.EvalSymlinks(string(self))
	if err != nil {
		return "", err
	}
	return FilePath(fp), nil
}

func (self FilePath) FromSlash() FilePath {
	return FilePath(filepath.FromSlash(string(self)))
}

func (self FilePath) Glob() ([]string, error) {
	return filepath.Glob(string(self))
}

func (self FilePath) IsLocal() bool {
	return filepath.IsLocal(string(self))
}

func (self FilePath) IsAbs() bool {
	return filepath.IsAbs(string(self))
}

func (self FilePath) Join(elem ...string) FilePath {
	elems := append([]string{string(self)}, elem...)
	return FilePath(filepath.Join(elems...))
}

func (self FilePath) Rel(dst FilePath) (FilePath, error) {
	fp, err := filepath.Rel(string(dst), string(self))
	if err != nil {
		return "", err
	}
	return FilePath(fp), nil
}

func (self FilePath) Match(pattern string) (bool, error) {
	return filepath.Match(pattern, string(self))
}

func (self FilePath) Split() (FilePath, string) {
	dir, file := filepath.Split(string(self))
	return FilePath(dir), file
}

func (self FilePath) SplitList() []string {
	return filepath.SplitList(string(self))
}

func (self FilePath) VolumeName() string {
	return filepath.VolumeName(string(self))
}

type WalkFunc func(path FilePath, info fs.FileInfo, err error) error

func (self FilePath) Walk(fn WalkFunc) error {
	return filepath.Walk(string(self), func(path string, info fs.FileInfo, err error) error {
		return fn(FilePath(path), info, err)
	})
}

type WalkDirFunc func(path FilePath, d fs.DirEntry, err error) error

func (self FilePath) WalkDir(fn WalkDirFunc) error {
	return filepath.WalkDir(string(self), func(path string, d fs.DirEntry, err error) error {
		return fn(FilePath(path), d, err)
	})
}
