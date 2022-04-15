package semantic

import (
	"GoParser/parser"
	"GoParser/parser/ast"
	"fmt"
)

func WaklTest(path string) {
	defer fff(path)
	tree := parser.ParseGoFile(path)
	(&Walker{}).WalkSourceFile(tree.(*ast.SourceFile))
}

func fff(path string) {
	if err := recover(); err != nil {
		fmt.Println("\n", path, err)
	}
}
