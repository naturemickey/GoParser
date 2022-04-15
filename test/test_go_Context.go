package main

import (
	"GoParser/parser"
	"GoParser/parser/ast"
	"GoParser/test/semantic"
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(4)
	go waklTest("./test/test_examples_go_context/test1.go")
	go waklTest("./test/test_examples_go_context/test2.go")
	go waklTest("./test/test_examples_go_context/test3.go")
	go waklTest("./test/test_examples_go_context/test4.go")
	wg.Wait()
}

func waklTest(path string) {
	defer wg.Done()
	defer fff(path)
	tree := parser.ParseGoFile(path)
	semantic.WalkSourceFile(tree.(*ast.SourceFile))
}

func fff(path string) {
	if err := recover(); err != nil {
		fmt.Println(path, err)
	}
}
