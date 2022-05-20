package main

import (
	"GoParser/parser"
	"GoParser/parser/antlr4"
	"GoParser/parser/ast"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	a := ParseGo("{\nif string(decrypt) != theMsg {\n\t\tt.Fatal(\"test fail! the msg is not equal\")\n\t}\n}")
	println(a.String())
}

func ParseGo(code string) ast.INode {
	input := antlr.NewInputStream(code)
	lexer := antlr4.NewGoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := antlr4.NewGoParser(stream)
	p.SetErrorHandler(antlr.NewBailErrorStrategy())
	// p.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	// p.BuildParseTrees = true
	tree := p.Block()

	return tree.Accept(&parser.GoParserVisitorImpl{}).(ast.INode)
}
