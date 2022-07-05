package os

import "os"

// IsExist 文件是否存在
func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

// OpenNewFile 打开新文件
func OpenNewFile(path string) (*os.File, error) {
	exist, err := IsExist(path)
	if err != nil {
		return nil, err
	}
	if exist {
		if err = os.Remove(path); err != nil {
			return nil, err
		}
	}
	return os.Create(path)
}
