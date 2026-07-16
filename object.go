package main

import (
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

type Object struct {
	Hash string
	Type string
	Data []byte
}

func NewBlob(path string) (*Object, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	header := fmt.Sprintf("blob %d%c", len(contents), 0)
	data := append([]byte(header), contents...)
	hash := sha1.Sum(data)
	return &Object{
		Hash: hex.EncodeToString(hash[:]),
		Type: "blob",
		Data: data,
	}, nil
}

// StoreObject to compress the Git object and store it in /objects
func StoreObject(obj *Object) error {
	dir := obj.Hash[:2]
	file := obj.Hash[2:]
	objectDir := filepath.Join(".goit", "objects", dir)
	objectPath := filepath.Join(objectDir, file)
	if err := os.MkdirAll(objectDir, 0755); err != nil {
		return err
	}
	if _, err := os.Stat(objectPath); err == nil {
		return nil
	}
	out, err := os.Create(objectPath)
	if err != nil {
		return err
	}
	defer out.Close()
	zw := zlib.NewWriter(out)
	defer zw.Close()
	_, err = zw.Write(obj.Data)
	if err != nil {
		return err
	}
	return nil
}
