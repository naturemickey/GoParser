package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/naturemickey/GoParser/parser"
	"github.com/naturemickey/GoParser/parser/antlr4"
	"github.com/naturemickey/GoParser/parser/ast"
	"github.com/naturemickey/GoParser/test/utils"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	utils.WalkDir("/Users/mickey/git/mis-backend/", parse_go_file, "/Users/mickey/git/mis-backend/vendor/")
	println(time.Since(start).Seconds())
}

func parse_go_file(filePath string) {
	if strings.HasSuffix(filePath, ".go") && !strings.HasSuffix(filePath, ".pb.go") {
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

		//_ = accept.String()
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
