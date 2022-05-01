package main

import (
	"GoParser/parser"
	"GoParser/parser/ast"
)

// packageName --> Annotations
var annotationMap = map[string]*Annotations{}

type Annotations struct {
	// methodName --> Annotation
	methodMap map[string]*Annotation
	// structName --> Annotation
	structMap map[string]*Annotation
	// fieldName --> Annotation, fieldName的构成中包括structName（例：sn.fn）
	fieldMap map[string]*Annotation
}

func NewAnnotations() *Annotations {
	return &Annotations{methodMap: map[string]*Annotation{}, structMap: map[string]*Annotation{}, fieldMap: map[string]*Annotation{}}
}

type Annotation struct {
}

func NewAnnotation(annotation string) *Annotation {
	return &Annotation{}
}

func main() {
	parseAnnotationGoFile("./test/testannotation/TestStru.go")

	println(annotationMap)
}

func parseAnnotationGoFile(filePath string) {
	sourceFile := parser.ParseGoFile(filePath).(*ast.SourceFile)
	pkgClau := sourceFile.PackageClause()
	packageName := pkgClau.PackageName()

	annotations := NewAnnotations()

	annotationMap[packageName] = annotations

	for _, fmd := range sourceFile.FunctionOrMethodOrDeclarations() {
		switch child := fmd.(type) {
		case *ast.MethodDecl:
			if child.Annotation() != "" {
				annotations.methodMap[child.FunName()] = NewAnnotation(child.Annotation())
			}
		case *ast.TypeDecl:
			for _, typeSpec := range child.TypeSpecs() {
				annotation := typeSpec.Annotation()
				name := typeSpec.Id()
				if annotation != "" {
					if typeSpec.Type_().TypeLit() == nil {
						panic("annotation 错误")
					}
					switch child := typeSpec.Type_().TypeLit().(type) {
					case *ast.StructType:
						annotations.structMap[name] = NewAnnotation(annotation)
						for _, fieldDecl := range child.FieldDecls() {
							fa := fieldDecl.Annotation()
							if fa != "" {
								idList := fieldDecl.IdentifierList()
								if idList != nil {
									if len(idList.Identifiers()) == 1 {
										fieldName := idList.Identifiers()[0]
										annotations.fieldMap[fieldName] = NewAnnotation(fa)
									} else {
										panic("annotation 错误")
									}
								}
							}
						}
					default:
						panic("annotation 错误")
					}
				}
			}
		}
	}
}
