package main

import (
	"fmt"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func TankaExtensions() []*jsonnet.NativeFunction {
	return []*jsonnet.NativeFunction{
		{
			Name:   "helloWorld",
			Params: ast.Identifiers{"author"},
			Func: func(data []interface{}) (interface{}, error) {
				return fmt.Sprintf("Hello world !! (by %s)", data[0]), nil
			},
		},
	}
}
