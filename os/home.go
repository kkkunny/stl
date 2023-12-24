package stlos

import "github.com/mitchellh/go-homedir"

// UserHomeDir 家目录地址
func UserHomeDir() (FilePath, error) {
	fp, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return NewFilePath(fp), nil
}
