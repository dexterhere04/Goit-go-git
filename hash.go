package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func hashObject(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	objectHeader := fmt.Sprintf("blob %d%c", stat.Size(), 0)
	var buffer [4096]byte
	hasher := sha1.New()
	_, err = hasher.Write([]byte(objectHeader))
	if err != nil {
		return err
	}
	for {
		n, err := file.Read(buffer[:])
		if n > 0 {
			_, writeErr := hasher.Write(buffer[:n])
			if writeErr != nil {
				return writeErr
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	hash := hasher.Sum(nil)
	fmt.Println(hex.EncodeToString(hash))
	return nil
}
