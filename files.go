package main

import (
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"strings"
)

// Method to get list of go files from the give parent
func getGoFiles(parent string, recursive, includeVendor bool) ([]string, error) {
	if verbose {
		color.Cyan("\n[Collecting .go files]")
	}
	var files []string
	err := filepath.Walk(parent,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && !recursive && info.Name() != strings.TrimSuffix(parent, "/") {
				return filepath.SkipDir
			}
			// Skipping vendor directory
			if info.IsDir() && info.Name() == "vendor" && !includeVendor {
				return filepath.SkipDir
			}
			// Only listing files
			if !info.IsDir() {
				splits := strings.Split(info.Name(), ".")
				if splits[len(splits)-1] == "go" {
					if verbose {
						color.White("- %s", path)
					}
					files = append(files, path)
				}
			}
			return nil
		})
	return files, err
}
