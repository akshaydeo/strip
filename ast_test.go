package main

import "testing"

func Test_commentFunctionCalls(t *testing.T) {
	commentFunctionCalls("demo/main.go", "log", "log.Println")
}

func Test_uncommentFunctionCalls(t *testing.T) {
	uncommentFunctionCalls("demo/main.go", "log", "log.Println")
}
