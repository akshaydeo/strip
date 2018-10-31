package main

import (
	"os"
	"path/filepath"
	"strings"
)

// Method to get list of go files from the give parent
func getGoFiles(parent string) ([]string, error) {
	var files []string
	err := filepath.Walk(parent,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Only listing files
			if !info.IsDir() {
				splits := strings.Split(info.Name(), ".")
				if splits[len(splits)-1] == "go" {
					files = append(files, info.Name())
				}
			}
			return nil
		})
	return files, err
}
