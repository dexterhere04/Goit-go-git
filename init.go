package main

import "os"

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
