package main

import (
	"fmt"
)

func hashObject(path string) error {
	obj, err := NewBlob(path)
	if err != nil {
		return err
	}

	fmt.Println(obj.Hash)

	return nil
}
