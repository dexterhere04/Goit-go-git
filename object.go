package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
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
