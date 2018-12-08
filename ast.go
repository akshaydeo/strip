package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

// Method to comment all function calls from given file
func commentFunctionCalls(fname, pkg, fun string) (bool, error) {
	if verbose {
		color.White("- %s", fname)
	}
	modified := false
	stripImport := true
	fileSet := token.NewFileSet() // positions are relative to fileSet
	node, err := parser.ParseFile(fileSet, fname, nil, parser.ParseComments)
	if err != nil {
		return false, err
	}

	//----------------------------------------------------------------------------------------------
	// Finding calls and commenting them out
	//----------------------------------------------------------------------------------------------
	for _, decl := range node.Decls {
		switch decl.(type) {
		case *ast.FuncDecl:
			fn, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}
			//color.Green("Checking function: %s", fn.Name.Name)
			// Iterating through the function statements
			for _, ln := range fn.Body.List {
				//fmt.Println(reflect.TypeOf(ln).String())
				switch ln.(type) {
				case *ast.ExprStmt:
					st := ln.(*ast.ExprStmt)
					//fmt.Println(reflect.TypeOf(st.X).String())
					switch st.X.(type) {
					case *ast.CallExpr:
						call := st.X.(*ast.CallExpr)
						selector := call.Fun.(*ast.SelectorExpr)
						finalCall := fmt.Sprintf("%s.%s", selector.X.(*ast.Ident).Name, call.Fun.(*ast.SelectorExpr).Sel.Name)
						if finalCall == fun {
							modified = true
							//fmt.Println("Got function call\n commenting it")
							selector.X.(*ast.Ident).Name = fmt.Sprintf("//%s", selector.X.(*ast.Ident).Name)
						} else {
							stripImport = false
						}
					}
				}
			}
		}

	}
	if stripImport {
		//----------------------------------------------------------------------------------------------
		// Removing imports
		//----------------------------------------------------------------------------------------------
		for _, imps := range node.Imports {
			if strings.Replace(imps.Path.Value, "\"", "", -1) == pkg {
				// Commenting out import
				modified = true
				imps.Path.Value = fmt.Sprintf("//%s", imps.Path.Value)
			}
		}
	}
	if !modified {
		return modified, nil
	}
	// Writing file
	// overwrite the file with modified version of ast.
	write, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer write.Close()
	w := bufio.NewWriter(write)
	err = format.Node(w, fileSet, node)
	if err != nil {
		panic(err)
	}
	w.Flush()
	return modified, nil
}

func uncommentFunctionCalls(fname, pkg, fun string) (bool, error) {
	if verbose {
		color.White("- %s", fname)
	}
	modified := false
	fileSet := token.NewFileSet() // positions are relative to fileSet
	node, err := parser.ParseFile(fileSet, fname, nil, parser.ParseComments)
	if err != nil {
		return false, err
	}
	//----------------------------------------------------------------------------------------------
	// Checking all the comments
	//----------------------------------------------------------------------------------------------
	for _, cmtg := range node.Comments {
		//fmt.Println(cmtg.Text())
		for _, cmt := range cmtg.List {
			ln := strings.Replace(cmt.Text, "//", "", 1)
			if strings.Replace(ln, "\"", "", -1) == pkg {
				modified = true
				cmt.Text = ln
			} else {
				// Check if this is one of the function call
				expr, err := parser.ParseExpr(ln)
				if err != nil {
					continue
				}
				switch expr.(type) {
				case *ast.CallExpr:
					call := expr.(*ast.CallExpr)
					selector := call.Fun.(*ast.SelectorExpr)
					finalCall := fmt.Sprintf("%s.%s", selector.X.(*ast.Ident).Name, call.Fun.(*ast.SelectorExpr).Sel.Name)
					if finalCall == fun {
						modified = true
						//fmt.Println("Got function call\n commenting it")
						cmt.Text = ln
					}

				}
			}
		}
	}
	if !modified {
		return modified, nil
	}
	// Writing file
	// overwrite the file with modified version of ast.
	write, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer write.Close()
	w := bufio.NewWriter(write)
	err = format.Node(w, fileSet, node)
	if err != nil {
		panic(err)
	}
	w.Flush()
	return modified, nil
}
