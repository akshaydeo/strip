package main

import (
	"flag"
	"github.com/fatih/color"
)

var verbose = false

func main() {
	//----------------------------------------------------------------------------------------------
	// Parsing args and validating
	//----------------------------------------------------------------------------------------------
	var calls arrayFlags
	pkg := flag.String("pkg", "", "Package associated with the calls")
	recursive := flag.Bool("r", true, "Perform recursively")
	path := flag.String("path", "./", "Path to perform changes")
	unStrip := flag.Bool("u", false, "Revert stripping")
	v := flag.Bool("v", true, "Verbose")
	verbose = *v
	flag.Var(&calls, "call", "Call you want to strip/unStrip")
	flag.Parse()
	//----------------------------------------------------------------------------------------------
	// Listing all the go files
	//----------------------------------------------------------------------------------------------
	if verbose {
		color.Green("---------------------------")
		if *unStrip {
			color.Green("UnStrip")
		} else {
			color.Green("Strip")
		}
		color.Green("---------------------------")
		color.Green("Calls: %s\nPackage: %s\nPath: %s\nRecursively: %v", calls.String(), *pkg, *path, *recursive)
	}
	//----------------------------------------------------------------------------------------------
	// Listing go files
	//----------------------------------------------------------------------------------------------
	files, err := getGoFiles(*path, false)
	if err != nil {
		panic(err)
	}
	//----------------------------------------------------------------------------------------------
	// Modified each file one by one
	//----------------------------------------------------------------------------------------------
	var modifiedFiles []string
	for _, file := range files {
		for _, call := range calls.Get() {
			if *unStrip {
				if verbose {
					color.Yellow("---------------------------")
					color.Yellow("UnStripping %s", call)
					color.Yellow("---------------------------")
				}
				modified, err := uncommentFunctionCalls(file, *pkg, call)
				if err != nil {
					panic(err)
				}
				if modified {
					modifiedFiles = append(modifiedFiles, file)
				}
				continue
			}
			if verbose {
				color.Yellow("---------------------------")
				color.Yellow("Stripping %s", call)
				color.Yellow("---------------------------")
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
	color.Red("------------------------")
	color.Red("Total modified files: %d", len(modifiedFiles))
	color.Red("------------------------")
	for _, file := range modifiedFiles {
		color.White("- %s", file)
	}
}
