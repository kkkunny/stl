package stlos

import "os"

// Exist 是否存在文件
func Exist(path FilePath) (bool, error) {
	_, err := os.Stat(string(path))
	if err != nil && os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
