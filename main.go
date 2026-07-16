package main

import (
	"fmt"
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
		write := false
		var filename string
		switch {
		case len(args) == 3:
			filename = args[2]

		case len(args) == 4 && args[2] == "-w":
			write = true
			filename = args[3]

		default:
			fmt.Println("Usage: goit hash-object [-w] <file>")
			return
		}
		err := hashObject(filename, write)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "cat-file":
		if len(args) != 4 {
			fmt.Println("Usage: goit cat-file <-p|-t|-s> <hash>")
			return
		}
		switch args[2] {
		case "-p":
			if err := catFile(args[3]); err != nil {
				fmt.Println(err)
			}
		case "-t":
			if err := catFileType(args[3]); err != nil {
				fmt.Println(err)
			}
		case "-s":
			if err := catFileSize(args[3]); err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("Usage: goit cat-file <-p|-t|-s> <hash>")
		}
	case "write-tree":
		hash, err := writeTree(".")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(hash)
	case "ls-tree":
		if len(args) != 3 {
			fmt.Println("Usage: goit ls-tree <hash>")
			return
		}
		if err := lsTree(args[2]); err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("unknown command")

	}
}
