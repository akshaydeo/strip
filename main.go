package main

import (
	"flag"
	"github.com/fatih/color"
)

var verbose = false
var recursive = false

func main() {
	//----------------------------------------------------------------------------------------------
	// Parsing args and validating
	//----------------------------------------------------------------------------------------------
	var calls arrayFlags
	pkg := flag.String("pkg", "", "Package associated with the calls")
	flag.BoolVar(&recursive, "r", false, "Perform recursively")
	path := flag.String("path", "./", "Path to perform changes")
	unStrip := flag.Bool("u", false, "Revert stripping")
	flag.BoolVar(&verbose, "v", false, "Verbose")
	flag.Var(&calls, "call", "Call you want to strip/unStrip")
	flag.Parse()
	//----------------------------------------------------------------------------------------------
	// Listing all the go files
	//----------------------------------------------------------------------------------------------
	if verbose {
		if *unStrip {
			color.Green("[UnStrip]")
		} else {
			color.Green("[Strip]")
		}
		color.White("Calls: %s\nPackage: %s\nPath: %s\nRecursively: %v", calls.String(), *pkg, *path, recursive)
	}
	//----------------------------------------------------------------------------------------------
	// Listing go files
	//----------------------------------------------------------------------------------------------
	files, err := getGoFiles(*path, recursive, false)
	if err != nil {
		panic(err)
	}
	//----------------------------------------------------------------------------------------------
	// Modified each file one by one
	//----------------------------------------------------------------------------------------------
	var modifiedFiles []string
	for _, call := range calls.Get() {
		if *unStrip {
			if verbose {
				color.Yellow("\n[UnStripping %s]", call)
			}
		} else {
			if verbose {
				color.Yellow("\n[Stripping %s]", call)
			}
		}
		for _, file := range files {
			if *unStrip {
				modified, err := uncommentFunctionCalls(file, *pkg, call)
				if err != nil {
					panic(err)
				}
				if modified {
					modifiedFiles = append(modifiedFiles, file)
				}
				continue
			}
			modified, err := commentFunctionCalls(file, *pkg, call)
			if err != nil {
				panic(err)
			}
			if modified {
				modifiedFiles = append(modifiedFiles, file)
			}
		}
	}
	color.Red("\n[Total modified files: %d]", len(modifiedFiles))
	for _, file := range modifiedFiles {
		color.White("- %s", file)
	}
}
