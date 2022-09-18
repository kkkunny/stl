package os

import (
	"os"
	"path/filepath"
	"strings"
)

// Path 路径
type Path string

func (self Path) String() string {
	return string(self)
}

// GetParent 获取父目录
func (self Path) GetParent() Path {
	return Path(filepath.Dir(string(self)))
}

// Join 拼接路径
func (self Path) Join(path Path) Path {
	return Path(filepath.Join(string(self), string(path)))
}

// Clean 清理路径
func (self Path) Clean() Path {
	return Path(filepath.Clean(string(self)))
}

// GetBase 获取最底层路径
func (self Path) GetBase() Path {
	return Path(filepath.Base(string(self)))
}

// GetExtension 获取拓展名
func (self Path) GetExtension() string {
	return filepath.Ext(string(self))
}

// IsAbsolute 是否是绝对路径
func (self Path) IsAbsolute() bool {
	return filepath.IsAbs(string(self))
}

// GetRelative 获取相对路径
func (self Path) GetRelative(base Path) (Path, error) {
	path, err := filepath.Rel(string(base), string(self))
	return Path(path), err
}

// IsExist 文件是否存在
func (self Path) IsExist() bool {
	exist, _ := IsExist(string(self))
	return exist
}

// GetAbsolute 获取绝对路径
func (self Path) GetAbsolute() (Path, error) {
	path, err := filepath.Abs(string(self))
	return Path(path), err
}

// WithExtension 替换拓展名
func (self Path) WithExtension(ext string) Path {
	index := strings.LastIndexByte(string(self), '.')
	if index >= 0 {
		return Path(string(self)[:index] + "." + ext)
	}
	return self
}

// IsDir 是否是目录
func (self Path) IsDir() bool {
	info, err := os.Stat(string(self))
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsFile 是否是文件
func (self Path) IsFile() bool {
	info, err := os.Stat(string(self))
	if err != nil {
		return false
	}
	return !info.IsDir()
}
