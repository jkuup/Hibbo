package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// 判断所给的路径是否是目录
func IsDir(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// 判断所给的路径是不是文件
func IsFile(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
	/*
		existed := true
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			existed = false
		}
		return existed
	*/
}

// 创建目录
func MakeDir(dir string) error {
	// 判断目录是否存在并创建
	if !IsDir(dir) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

// 删除目录
func RemoveDir(dir string) error {
	// 判断目录是否存在在删除
	if !IsDir(dir) {
		return errors.New("cannot delete without directory")
	}

	return os.RemoveAll(dir)
}

// 拷贝文件
func CopyFile(src, dst string) (int64, error) {
	// 判断文件是否存在
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	// 判断文件是否为常规文件，是否可以打开
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	// 打开文件
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	// 拷贝文件
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
