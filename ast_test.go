package main

import "testing"

func Test_commentFunctionCalls(t *testing.T) {
	commentFunctionCalls("demo/main.go","log.Println")
}
