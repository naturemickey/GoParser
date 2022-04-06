package main

import (
	"GoParser/parser"
	"GoParser/parser/antlr4"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
)

func main() {
	file, _ := ioutil.ReadFile("./test_examples/anonymousMethods.go")
	input := antlr.NewInputStream(string(file))
	lexer := antlr4.NewGoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := antlr4.NewGoParser(stream)
	p.SetErrorHandler(antlr.NewBailErrorStrategy())
	// p.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	// p.BuildParseTrees = true
	tree := p.SourceFile()

	visitor := &parser.GoParserVisitorImpl{}

	accept := tree.Accept(visitor)

	println(accept)
}
