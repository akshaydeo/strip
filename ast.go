package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// Method to comment all function calls from given file
func commentFunctionCalls(fname, fun string) error {
	fileSet := token.NewFileSet() // positions are relative to fileSet
	node, err := parser.ParseFile(fileSet, fname, nil, 0)
	if err != nil {
		return err
	}
	for _, decl := range node.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		fmt.Println("Checking function =>", fn.Name.Name)
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
						fmt.Println("Got function call")

					}
				}
			}
		}
	}
	return nil
}

func uncommentFunctionCalls(fname, fun string) error {
	return nil
}
