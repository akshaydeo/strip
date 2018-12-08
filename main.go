package main

import (
	"flag"
	"github.com/fatih/color"
)

func main() {
	//----------------------------------------------------------------------------------------------
	// Parsing args and validating
	//----------------------------------------------------------------------------------------------
	var calls arrayFlags
	pkg := flag.String("pkg", "", "Package associated with the calls")
	recursive := flag.Bool("r", true, "Perform recursively")
	path := flag.String("path", "./", "Path to perform changes")
	unStrip := flag.Bool("u", false, "Revert stripping")
	flag.Var(&calls, "call", "Call you want to strip/unStrip")
	flag.Parse()
	//----------------------------------------------------------------------------------------------
	// Listing all the go files
	//----------------------------------------------------------------------------------------------
	color.Green("---------------------------")
	if *unStrip {
		color.Green("UnStrip")
	} else {
		color.Green("Strip")
	}
	color.Green("---------------------------")
	color.Green("Calls: %s\nPackage: %s\nPath: %s\nRecursively: %v", calls.String(), *pkg, *path, *recursive)
	color.Green("---------------------------")
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
	for _, file := range files {
		for _, call := range calls.Get() {
			if *unStrip {
				color.Yellow("UnStripping %s", call)
				err = uncommentFunctionCalls(file, *pkg, call)
				if err != nil {
					panic(err)
				}
				continue
			}
			color.Yellow("Stripping %s", call)
			err = commentFunctionCalls(file, *pkg, call)
			if err != nil {
				panic(err)
			}
		}

	}
}
