package main

import (
	"GoParser/parser"
	"GoParser/parser/antlr4"
	"GoParser/parser/ast"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
	"strings"
)

func main() {
	parse_dir("/Users/mickey/git/mis-backend/")
}

func parse_dir(path string) {
	dir, err2 := ioutil.ReadDir(path)
	if err2 != nil {
		println(err2.Error())
	}
	for _, fileInfo := range dir {
		if fileInfo.IsDir() {
			parse_dir(path + fileInfo.Name() + "/")
		} else {
			parse_go_file(path + fileInfo.Name())
		}
	}
}

func parse_go_file(filePath string) {
	if strings.HasSuffix(filePath, ".go") {
		//println("=======================================================================================")
		println("开始处理文件：", filePath)

		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err.Error())
		}
		originContent := string(file)

		//println("原始文件：\n", originContent)

		input := antlr.NewInputStream(originContent)
		lexer := antlr4.NewGoLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := antlr4.NewGoParser(stream)
		p.SetErrorHandler(antlr.NewBailErrorStrategy())
		// p.SetErrorHandler(antlr.NewDefaultErrorStrategy())
		// p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
		// p.BuildParseTrees = true
		tree := p.SourceFile()

		visitor := &parser.GoParserVisitorImpl{}

		accept := tree.Accept(visitor).(ast.INode)

		content := accept.String()

		//println("------------------------")
		//println("解析后：\n", content)

		//bs := ast.GoFmt(content)
		_ = ast.GoFmt(content)

		//println("------------------------")
		//println("gofmt后：\n", string(bs))
		//println("=======================================================================================")
	}
}
