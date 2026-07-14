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
		err := initRepo()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Initialized repo successfully")
	default:
		fmt.Println("unknown command")

	}
}

func initRepo() error {
	if err := createDirs(); err != nil {
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
