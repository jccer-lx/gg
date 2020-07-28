package helper

import "os"

//判断目录
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//创建目录
func CreateDir(path string) error {
	ok, _ := PathExists(path)
	if ok {
		return nil
	}
	return os.Mkdir(path, 0755)
}
