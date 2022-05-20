package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/naturemickey/GoParser/parser"
	"github.com/naturemickey/GoParser/parser/antlr4"
	"github.com/naturemickey/GoParser/parser/ast"
	"io/ioutil"
	"strings"
)

func main() {
	path := "./test_examples/"
	dir, err2 := ioutil.ReadDir(path)
	if err2 != nil {
		println(err2.Error())
	}
	whiteList := []string{}
	for _, fileInfo := range dir {
		if strings.HasSuffix(fileInfo.Name(), "_go") {

			if len(whiteList) > 0 {
				var in bool = false
				for _, s := range whiteList {
					if s == fileInfo.Name() {
						in = true
						break
					}
				}
				if !in {
					continue
				}
			}

			println("=======================================================================================")
			filePath := path + fileInfo.Name()

			println("开始处理文件：", filePath)

			file, err := ioutil.ReadFile(filePath)
			if err != nil {
				panic(err.Error())
			}
			originContent := string(file)

			println("原始文件：\n", originContent)

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

			println("------------------------")
			println("解析后：\n", content)

			bs := ast.GoFmt(content)

			println("------------------------")
			println("gofmt后：\n", string(bs))
			println("=======================================================================================")
		}
	}
}
