package parser

import (
	"GoParser/parser/antlr4"
	"GoParser/parser/ast"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
)

var visitor = &GoParserVisitorImpl{}

func ParseGoFile(fileName string) ast.INode {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	code := string(file)
	return ParseGoCode(code)
}

func ParseGoCode(code string) ast.INode {
	input := antlr.NewInputStream(code)
	lexer := antlr4.NewGoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := antlr4.NewGoParser(stream)
	p.SetErrorHandler(antlr.NewBailErrorStrategy())
	// p.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	// p.BuildParseTrees = true
	tree := p.SourceFile()

	return tree.Accept(visitor).(ast.INode)
}
