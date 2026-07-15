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
