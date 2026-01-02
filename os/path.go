package stlos

import (
	"path/filepath"
)

// ReplaceBase 替换path的from基路径为to
func ReplaceBase(path, from, to string) (string, error) {
	relFp, err := filepath.Rel(from, path)
	if err != nil {
		return "", err
	}
	return filepath.Join(to, relFp), nil
}
