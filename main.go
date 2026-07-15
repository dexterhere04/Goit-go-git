package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args[1:]) == 0 {
		fmt.Println("Really?,use a argument dude : goit <command>")
		return
	}
	command := args[1]
	switch command {
	case "init":
		var err error
		const DefBranch = "main"
		if len(args) < 3 {
			err = initRepo(DefBranch)
		} else {
			err = initRepo(args[2])
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Initialized repo successfully")
	case "hash-object":
		if len(args) < 3 {
			fmt.Println("no file name provided,Usage: goit hash-object <filename>")
			return
		}
		if err := hashObject(args[2]); err != nil {
			fmt.Println(err)
			return
		}

	default:
		fmt.Println("unknown command")

	}
}

func initRepo(defBranch string) error {
	if err := createDirs(); err != nil {
		return err
	}
	// head points to the which branch is currently checked out
	// refs/heads/master or heads/main contains SHA1 of the latest commit on branch
	if err := createHead(defBranch); err != nil {
		return err
	}
	return nil
}
func createDirs() error {
	paths := []string{".goit/objects", ".goit/refs/heads", ".goit/refs/tags"}
	for i := 0; i < len(paths); i++ {
		if err := os.MkdirAll(paths[i], 0755); err != nil {
			return err
		}
	}
	return nil
}
func createHead(defBranch string) error {
	file, err := os.Create(".goit/HEAD")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString("ref: refs/heads/" + defBranch)
	if err != nil {
		return err
	}
	return nil
}
func hashObject(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	var buffer [4096]byte
	hasher := sha1.New()
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
