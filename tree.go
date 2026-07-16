package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type TreeEntry struct {
	Mode string
	Name string
	Hash string
}

func writeTree(dir string) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	var treeEntries []TreeEntry

	for _, entry := range entries {
		name := entry.Name()
		if name == ".goit" {
			continue
		}
		path := filepath.Join(dir, name)

		if entry.IsDir() {
			subtreeHash, err := writeTree(path)
			if err != nil {
				return "", err
			}
			treeEntries = append(treeEntries, TreeEntry{
				Mode: "40000",
				Name: name,
				Hash: subtreeHash,
			})
		} else {
			blob, err := NewBlob(path)
			if err != nil {
				return "", err
			}
			if err := StoreObject(blob); err != nil {
				return "", err
			}
			treeEntries = append(treeEntries, TreeEntry{
				Mode: "100644",
				Name: name,
				Hash: blob.Hash,
			})
		}
	}

	sort.Slice(treeEntries, func(i, j int) bool {
		return treeEntries[i].Name < treeEntries[j].Name
	})

	var body []byte
	for _, e := range treeEntries {
		body = append(body, []byte(e.Mode+" "+e.Name)...)
		body = append(body, 0)
		hashBytes, err := hex.DecodeString(e.Hash)
		if err != nil {
			return "", err
		}
		body = append(body, hashBytes...)
	}

	header := fmt.Sprintf("tree %d%c", len(body), 0)
	data := append([]byte(header), body...)
	hash := sha1.Sum(data)
	hashStr := hex.EncodeToString(hash[:])

	treeObj := &Object{
		Hash: hashStr,
		Type: "tree",
		Data: data,
	}
	if err := StoreObject(treeObj); err != nil {
		return "", err
	}

	return hashStr, nil
}

func readTree(body []byte) []TreeEntry {
	var entries []TreeEntry
	for len(body) > 0 {
		spaceIdx := 0
		for body[spaceIdx] != ' ' {
			spaceIdx++
		}
		mode := string(body[:spaceIdx])
		body = body[spaceIdx+1:]

		nullIdx := 0
		for body[nullIdx] != 0 {
			nullIdx++
		}
		name := string(body[:nullIdx])
		body = body[nullIdx+1:]

		hash := hex.EncodeToString(body[:20])
		body = body[20:]

		entries = append(entries, TreeEntry{
			Mode: mode,
			Name: name,
			Hash: hash,
		})
	}
	return entries
}
