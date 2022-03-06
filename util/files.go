package util

import (
	"errors"
	"os"
)

func MakeDirectory(path string) error {
	if len(path) == 0 || path == "" {
		return errors.New("file path is empty")
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	return nil
}

func RemoveDriectory(path string) error {
	if len(path) == 0 || path == "" {
		return errors.New("file path is empty")
	}

	err := os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}

func RemoveFile(path string) error {
	if len(path) == 0 || path == "" {
		return errors.New("file path is empty")
	}

	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func CheckExistFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func CheckExistFileWithSize(path string) bool {
	fi, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	if fi.Size() > 0 {
		return true
	}

	return false
}
