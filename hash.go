package main

import (
	"fmt"
)

func hashObject(path string, write bool) error {
	obj, err := NewBlob(path)
	if err != nil {
		return err
	}
	if write {
		if err := StoreObject(obj); err != nil {
			return err
		}
	}
	fmt.Println(obj.Hash)
	return nil
}
func catFile(hash string) error {
	obj, err := ReadObject(hash)
	if err != nil {
		return err
	}
	fmt.Print(string(obj.Data))
	return nil
}

func catFileType(hash string) error {
	obj, err := ReadObject(hash)
	if err != nil {
		return err
	}
	fmt.Println(obj.Type)
	return nil
}

func catFileSize(hash string) error {
	obj, err := ReadObject(hash)
	if err != nil {
		return err
	}
	fmt.Println(len(obj.Data))
	return nil
}
