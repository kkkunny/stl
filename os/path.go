package stlos

import (
	"io/fs"
	"os"
	"path/filepath"
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

func (fp FilePath) Abs() (FilePath, error) {
	absFp, err := filepath.Abs(string(fp))
	if err != nil {
		return "", err
	}
	return FilePath(absFp), nil
}

func (fp FilePath) Base() string {
	return filepath.Base(string(fp))
}

func (fp FilePath) Clean() FilePath {
	return FilePath(filepath.Clean(string(fp)))
}

func (fp FilePath) Dir() FilePath {
	return FilePath(filepath.Dir(string(fp)))
}

func (fp FilePath) Ext() string {
	return filepath.Ext(string(fp))
}

func (fp FilePath) EvalSymlinks() (FilePath, error) {
	symFp, err := filepath.EvalSymlinks(string(fp))
	if err != nil {
		return "", err
	}
	return FilePath(symFp), nil
}

func (fp FilePath) FromSlash() FilePath {
	return FilePath(filepath.FromSlash(string(fp)))
}

func (fp FilePath) Glob() ([]string, error) {
	return filepath.Glob(string(fp))
}

func (fp FilePath) IsLocal() bool {
	return filepath.IsLocal(string(fp))
}

func (fp FilePath) IsAbs() bool {
	return filepath.IsAbs(string(fp))
}

func (fp FilePath) Join(elem ...string) FilePath {
	elems := append([]string{string(fp)}, elem...)
	return FilePath(filepath.Join(elems...))
}

func (fp FilePath) Rel(dst FilePath) (string, error) {
	relFp, err := filepath.Rel(string(dst), string(fp))
	if err != nil {
		return "", err
	}
	return relFp, nil
}

func (fp FilePath) Match(pattern string) (bool, error) {
	return filepath.Match(pattern, string(fp))
}

func (fp FilePath) Split() (FilePath, string) {
	dir, file := filepath.Split(string(fp))
	return FilePath(dir), file
}

func (fp FilePath) SplitList() []string {
	return filepath.SplitList(string(fp))
}

func (fp FilePath) VolumeName() string {
	return filepath.VolumeName(string(fp))
}

func (fp FilePath) ReplaceBase(from, to FilePath) (FilePath, error) {
	relFp, err := fp.Rel(from)
	if err != nil {
		return "", err
	}
	return to.Join(relFp), nil
}

type WalkFunc func(path FilePath, info fs.FileInfo, err error) error

func (fp FilePath) Walk(fn WalkFunc) error {
	return filepath.Walk(string(fp), func(path string, info fs.FileInfo, err error) error {
		return fn(FilePath(path), info, err)
	})
}

type WalkDirFunc func(path FilePath, d fs.DirEntry, err error) error

func (fp FilePath) WalkDir(fn WalkDirFunc) error {
	return filepath.WalkDir(string(fp), func(path string, d fs.DirEntry, err error) error {
		return fn(FilePath(path), d, err)
	})
}
