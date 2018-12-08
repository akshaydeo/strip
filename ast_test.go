package main

import "testing"

func Test_commentFunctionCalls(t *testing.T) {
	commentFunctionCalls("demo/main.go", "github.com/Sirupsen/logrus", "logrus.Debug")
}

func Test_uncommentFunctionCalls(t *testing.T) {
	uncommentFunctionCalls("demo/main.go", "github.com/Sirupsen/logrus", "logrus.Debug")
}
