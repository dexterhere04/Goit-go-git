package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"
)

func commitTree(treeHash, parent, message string) (string, error) {
	now := time.Now()
	timestamp := now.Unix()
	_, offset := now.Zone()
	tz := fmt.Sprintf("%+05d", offset/3600*100+offset%3600/60)

	author := "goit <goit@local>"
	committer := author

	content := fmt.Sprintf("tree %s\n", treeHash)
	if parent != "" {
		content += fmt.Sprintf("parent %s\n", parent)
	}
	content += fmt.Sprintf("author %s %d %s\n", author, timestamp, tz)
	content += fmt.Sprintf("committer %s %d %s\n", committer, timestamp, tz)
	content += "\n"
	content += message + "\n"

	header := fmt.Sprintf("commit %d%c", len(content), 0)
	data := append([]byte(header), []byte(content)...)
	hash := sha1.Sum(data)
	hashStr := hex.EncodeToString(hash[:])

	obj := &Object{
		Hash: hashStr,
		Type: "commit",
		Data: data,
	}
	if err := StoreObject(obj); err != nil {
		return "", err
	}
	return hashStr, nil
}

func readHeadBranch() (string, error) {
	data, err := os.ReadFile(".goit/HEAD")
	if err != nil {
		return "", err
	}
	ref := string(data)
	if len(ref) < 6 || ref[:5] != "ref: " {
		return "", fmt.Errorf("invalid HEAD")
	}
	branchRef := ref[5:]
	if branchRef[len(branchRef)-1] == '\n' {
		branchRef = branchRef[:len(branchRef)-1]
	}
	return branchRef, nil
}

func readRef(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := string(data)
	if hash[len(hash)-1] == '\n' {
		hash = hash[:len(hash)-1]
	}
	return hash, nil
}

func writeRef(path, hash string) error {
	fullPath := ".goit/" + path
	dir := fullPath[:strings.LastIndex(fullPath, "/")]
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(fullPath, []byte(hash+"\n"), 0644)
}

func goitCommit(message string) (string, error) {
	treeHash, err := writeTree(".")
	if err != nil {
		return "", err
	}

	var parent string
	branchRef, err := readHeadBranch()
	if err == nil {
		if parentHash, err := readRef(".goit/" + branchRef); err == nil {
			parent = parentHash
		}
	}

	commitHash, err := commitTree(treeHash, parent, message)
	if err != nil {
		return "", err
	}

	if branchRef, err := readHeadBranch(); err == nil {
		writeRef(branchRef, commitHash)
	}

	return commitHash, nil
}
