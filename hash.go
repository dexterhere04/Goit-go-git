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
