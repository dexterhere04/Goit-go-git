package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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
func ReadObject(hash string) (*Object, error) {

	dir := hash[:2]
	file := hash[2:]

	objectPath := filepath.Join(".goit", "objects", dir, file)

	f, err := os.Open(objectPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Decompress using zlib.
	zr, err := zlib.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	// Read the complete decompressed object.
	data, err := io.ReadAll(zr)
	if err != nil {
		return nil, err
	}

	// Find the NULL byte separating header and body.
	idx := bytes.IndexByte(data, 0)
	if idx == -1 {
		return nil, fmt.Errorf("invalid git object")
	}

	header := string(data[:idx])
	body := data[idx+1:]

	// Header format:
	// blob 123
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid object header")
	}

	return &Object{
		Hash: hash,
		Type: parts[0],
		Data: body,
	}, nil
}
