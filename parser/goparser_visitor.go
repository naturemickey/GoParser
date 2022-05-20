package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/naturemickey/GoParser/parser/antlr4"
	"github.com/naturemickey/GoParser/parser/ast"
)

type GoParserVisitorImpl struct {
	antlr.BaseParseTreeVisitor
}

func (visitor *GoParserVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(visitor)
}
func (visitor *GoParserVisitorImpl) VisitChildren(node antlr.RuleNode) interface{} {
	panic("implement me")
}
func (visitor *GoParserVisitorImpl) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("implement me")
}
func (visitor *GoParserVisitorImpl) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitSourceFile(ctx *antlr4.SourceFileContext) interface{} {
	packageClause := ctx.PackageClause().Accept(visitor).(*ast.PackageClause)
	importDecls := []*ast.ImportDecl{}
	for _, decl := range ctx.AllImportDecl() {
		importDecls = append(importDecls, decl.Accept(visitor).(*ast.ImportDecl))
	}
	functionOrMethodOrDeclarations := []ast.IFunctionMethodDeclaration{}
	for _, tree := range ctx.GetChildren() {
		switch child := tree.(type) {
		case *antlr4.FunctionDeclContext:
			functionOrMethodOrDeclarations = append(functionOrMethodOrDeclarations, child.Accept(visitor).(*ast.FunctionDecl))
		case *antlr4.MethodDeclContext:
			functionOrMethodOrDeclarations = append(functionOrMethodOrDeclarations, child.Accept(visitor).(*ast.MethodDecl))
		case *antlr4.DeclarationContext:
			functionOrMethodOrDeclarations = append(functionOrMethodOrDeclarations, child.Accept(visitor).(ast.Declaration))
		}
	}
	return ast.NewSourceFile(packageClause, importDecls, functionOrMethodOrDeclarations)
}

func (visitor *GoParserVisitorImpl) VisitPackageClause(ctx *antlr4.PackageClauseContext) interface{} {
	packageName := ctx.GetPackageName().GetText()
	return ast.NewPackageClause(packageName)
}

func (visitor *GoParserVisitorImpl) VisitImportDecl(ctx *antlr4.ImportDeclContext) interface{} {
	importSpecs := []*ast.ImportSpec{}
	for _, spec := range ctx.AllImportSpec() {
		importSpecs = append(importSpecs, spec.Accept(visitor).(*ast.ImportSpec))
	}
	return ast.NewImportDecl(importSpecs)
}

func (visitor *GoParserVisitorImpl) VisitImportSpec(ctx *antlr4.ImportSpecContext) interface{} {
	var alias string
	if ctx.GetAlias() != nil {
		alias = ctx.GetAlias().GetText()
	}
	importPath := ctx.ImportPath().Accept(visitor).(*ast.ImportPath)
	return ast.NewImportSpec(alias, importPath)
}

func (visitor *GoParserVisitorImpl) VisitImportPath(ctx *antlr4.ImportPathContext) interface{} {
	path := ctx.GetText()
	return ast.NewImportPath(path)
}

func (visitor *GoParserVisitorImpl) VisitDeclaration(ctx *antlr4.DeclarationContext) interface{} {
	if ctx.ConstDecl() != nil {
		return ctx.ConstDecl().Accept(visitor)
	}
	if ctx.TypeDecl() != nil {
		return ctx.TypeDecl().Accept(visitor)
	}
	return ctx.VarDecl().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitConstDecl(ctx *antlr4.ConstDeclContext) interface{} {
	var constSpecs []*ast.ConstSpec
	for _, spec := range ctx.AllConstSpec() {
		constSpecs = append(constSpecs, spec.Accept(visitor).(*ast.ConstSpec))
	}
	return ast.NewConstDecl(constSpecs)
}

func (visitor *GoParserVisitorImpl) VisitConstSpec(ctx *antlr4.ConstSpecContext) interface{} {
	var identifierList *ast.IdentifierList
	var type_ *ast.Type_
	var expressionList *ast.ExpressionList

	identifierList = ctx.IdentifierList().Accept(visitor).(*ast.IdentifierList)
	if ctx.Type_() != nil {
		type_ = ctx.Type_().Accept(visitor).(*ast.Type_)
	}
	if ctx.ExpressionList() != nil {
		expressionList = ctx.ExpressionList().Accept(visitor).(*ast.ExpressionList)
	}
	return ast.NewConstSpec(identifierList, type_, expressionList)
}

func (visitor *GoParserVisitorImpl) VisitIdentifierList(ctx *antlr4.IdentifierListContext) interface{} {
	var identifiers []string
	for _, id := range ctx.AllIDENTIFIER() {
		identifiers = append(identifiers, id.GetText())
	}
	return ast.NewIdentifierList(identifiers)
}

func (visitor *GoParserVisitorImpl) VisitExpressionList(ctx *antlr4.ExpressionListContext) interface{} {
	var expressions []*ast.Expression
	for _, expression := range ctx.AllExpression() {
		expressions = append(expressions, expression.Accept(visitor).(*ast.Expression))
	}
	return ast.NewExpressionList(expressions)
}

func (visitor *GoParserVisitorImpl) VisitTypeDecl(ctx *antlr4.TypeDeclContext) interface{} {
	var typeSpecs []*ast.TypeSpec
	for _, spec := range ctx.AllTypeSpec() {
		typeSpecs = append(typeSpecs, spec.Accept(visitor).(*ast.TypeSpec))
	}
	return ast.NewTypeDecl(typeSpecs)
}

func (visitor *GoParserVisitorImpl) VisitTypeSpec(ctx *antlr4.TypeSpecContext) interface{} {
	var annotation string
	//if ctx.ANNOTATION() != nil {
	//	annotation = ctx.ANNOTATION().GetText()
	//}
	id := ctx.IDENTIFIER().GetText()
	type_ := ctx.Type_().Accept(visitor).(*ast.Type_)
	return ast.NewTypeSpec(annotation, id, type_)
}

func (visitor *GoParserVisitorImpl) VisitFunctionDecl(ctx *antlr4.FunctionDeclContext) interface{} {
	var funName string = ctx.IDENTIFIER().GetText()
	var signature *ast.Signature = ctx.Signature().Accept(visitor).(*ast.Signature)
	var block *ast.Block
	if ctx.Block() != nil {
		block = ctx.Block().Accept(visitor).(*ast.Block)
	}
	return ast.NewFunctionDecl(funName, signature, block)
}

func (visitor *GoParserVisitorImpl) VisitMethodDecl(ctx *antlr4.MethodDeclContext) interface{} {
	var annotation string
	//if ctx.ANNOTATION() != nil {
	//	annotation = ctx.ANNOTATION().GetText()
	//}
	var receiver ast.Receiver = ctx.Receiver().Accept(visitor).(ast.Receiver)
	var funName string = ctx.IDENTIFIER().GetText()
	var signature *ast.Signature = ctx.Signature().Accept(visitor).(*ast.Signature)
	var block *ast.Block
	if ctx.Block() != nil {
		block = ctx.Block().Accept(visitor).(*ast.Block)
	}
	return ast.NewMethodDecl(annotation, receiver, funName, signature, block)
}

func (visitor *GoParserVisitorImpl) VisitReceiver(ctx *antlr4.ReceiverContext) interface{} {
	return ctx.Parameters().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitVarDecl(ctx *antlr4.VarDeclContext) interface{} {
	var varSpecs []*ast.VarSpec
	for _, spec := range ctx.AllVarSpec() {
		varSpecs = append(varSpecs, spec.Accept(visitor).(*ast.VarSpec))
	}
	return ast.NewVarDecl(varSpecs)
}

func (visitor *GoParserVisitorImpl) VisitVarSpec(ctx *antlr4.VarSpecContext) interface{} {
	var identifierList *ast.IdentifierList = ctx.IdentifierList().Accept(visitor).(*ast.IdentifierList)
	var type_ *ast.Type_
	var expressionList *ast.ExpressionList
	if ctx.Type_() != nil {
		type_ = ctx.Type_().Accept(visitor).(*ast.Type_)
	}
	if ctx.ExpressionList() != nil {
		expressionList = ctx.ExpressionList().Accept(visitor).(*ast.ExpressionList)
	}
	return ast.NewVarSpec(identifierList, type_, expressionList)
}

func (visitor *GoParserVisitorImpl) VisitBlock(ctx *antlr4.BlockContext) interface{} {
	var statementList *ast.StatementList
	if ctx.StatementList() != nil {
		statementList = ctx.StatementList().Accept(visitor).(*ast.StatementList)
	}
	return ast.NewBlock(statementList)
}

func (visitor *GoParserVisitorImpl) VisitStatementList(ctx *antlr4.StatementListContext) interface{} {
	var statements []ast.Statement
	for _, statement := range ctx.AllStatement() {
		statements = append(statements, statement.Accept(visitor).(ast.Statement))
	}
	return ast.NewStatementList(statements)
}

func (visitor *GoParserVisitorImpl) VisitStatement(ctx *antlr4.StatementContext) interface{} {
	if ctx.Declaration() != nil {
		return ctx.Declaration().Accept(visitor)
	}
	if ctx.LabeledStmt() != nil {
		return ctx.LabeledStmt().Accept(visitor)
	}
	if ctx.SimpleStmt() != nil {
		return ctx.SimpleStmt().Accept(visitor)
	}
	if ctx.GoStmt() != nil {
		return ctx.GoStmt().Accept(visitor)
	}
	if ctx.ReturnStmt() != nil {
		return ctx.ReturnStmt().Accept(visitor)
	}
	if ctx.BreakStmt() != nil {
		return ctx.BreakStmt().Accept(visitor)
	}
	if ctx.ContinueStmt() != nil {
		return ctx.ContinueStmt().Accept(visitor)
	}
	if ctx.GotoStmt() != nil {
		return ctx.GotoStmt().Accept(visitor)
	}
	if ctx.FallthroughStmt() != nil {
		return ctx.FallthroughStmt().Accept(visitor)
	}
	if ctx.Block() != nil {
		return ctx.Block().Accept(visitor)
	}
	if ctx.IfStmt() != nil {
		return ctx.IfStmt().Accept(visitor)
	}
	if ctx.SwitchStmt() != nil {
		return ctx.SwitchStmt().Accept(visitor)
	}
	if ctx.SelectStmt() != nil {
		return ctx.SelectStmt().Accept(visitor)
	}
	if ctx.ForStmt() != nil {
		return ctx.ForStmt().Accept(visitor)
	}
	return ctx.DeferStmt().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitSimpleStmt(ctx *antlr4.SimpleStmtContext) interface{} {
	if ctx.SendStmt() != nil {
		return ctx.SendStmt().Accept(visitor)
	}
	if ctx.IncDecStmt() != nil {
		return ctx.IncDecStmt().Accept(visitor)
	}
	if ctx.Assignment() != nil {
		return ctx.Assignment().Accept(visitor)
	}
	if ctx.ExpressionStmt() != nil {
		return ctx.ExpressionStmt().Accept(visitor)
	}
	return ctx.ShortVarDecl().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitExpressionStmt(ctx *antlr4.ExpressionStmtContext) interface{} {
	return ctx.Expression().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitSendStmt(ctx *antlr4.SendStmtContext) interface{} {
	var channel *ast.Expression = ctx.GetChannel().Accept(visitor).(*ast.Expression)
	var expression *ast.Expression = ctx.Expression(1).Accept(visitor).(*ast.Expression)
	return ast.NewSendStmt(channel, expression)
}

func (visitor *GoParserVisitorImpl) VisitIncDecStmt(ctx *antlr4.IncDecStmtContext) interface{} {
	expression := ctx.Expression().Accept(visitor).(*ast.Expression)
	isPlus := ctx.PLUS_PLUS() != nil
	return ast.NewIncDecStmt(expression, isPlus)
}

func (visitor *GoParserVisitorImpl) VisitAssignment(ctx *antlr4.AssignmentContext) interface{} {
	expressionListLeft := ctx.ExpressionList(0).Accept(visitor).(*ast.ExpressionList)
	assignOp := ctx.Assign_op().Accept(visitor).(*ast.AssignOp)
	expressionListRight := ctx.ExpressionList(1).Accept(visitor).(*ast.ExpressionList)
	return ast.NewAssignment(expressionListLeft, assignOp, expressionListRight)
}

func (visitor *GoParserVisitorImpl) VisitAssign_op(ctx *antlr4.Assign_opContext) interface{} {
	var prefix string
	if ctx.GetPrefix() != nil {
		prefix = ctx.GetPrefix().GetText()
	}
	return ast.NewAssignOp(prefix)
}

func (visitor *GoParserVisitorImpl) VisitShortVarDecl(ctx *antlr4.ShortVarDeclContext) interface{} {
	identifierList := ctx.IdentifierList().Accept(visitor).(*ast.IdentifierList)
	expressionList := ctx.ExpressionList().Accept(visitor).(*ast.ExpressionList)
	return ast.NewShortValDecl(identifierList, expressionList)
}

func (visitor *GoParserVisitorImpl) VisitEmptyStmt(ctx *antlr4.EmptyStmtContext) interface{} {
	return ast.NewEmptStmt()
}

func (visitor *GoParserVisitorImpl) VisitLabeledStmt(ctx *antlr4.LabeledStmtContext) interface{} {
	label := ctx.IDENTIFIER().GetText()
	var statement ast.Statement
	if ctx.Statement() != nil {
		statement = ctx.Statement().Accept(visitor).(ast.Statement)
	}
	return ast.NewLabeledStmt(label, statement)
}

func (visitor *GoParserVisitorImpl) VisitReturnStmt(ctx *antlr4.ReturnStmtContext) interface{} {
	var expressionList *ast.ExpressionList
	if ctx.ExpressionList() != nil {
		expressionList = ctx.ExpressionList().Accept(visitor).(*ast.ExpressionList)
	}
	return ast.NewReturnStmt(expressionList)
}

func (visitor *GoParserVisitorImpl) VisitBreakStmt(ctx *antlr4.BreakStmtContext) interface{} {
	var label string
	if ctx.IDENTIFIER() != nil {
		label = ctx.IDENTIFIER().GetText()
	}
	return ast.NewBreakStmt(label)
}

func (visitor *GoParserVisitorImpl) VisitContinueStmt(ctx *antlr4.ContinueStmtContext) interface{} {
	var label string
	if ctx.IDENTIFIER() != nil {
		label = ctx.IDENTIFIER().GetText()
	}
	return ast.NewContinueStmt(label)
}

func (visitor *GoParserVisitorImpl) VisitGotoStmt(ctx *antlr4.GotoStmtContext) interface{} {
	var label string
	if ctx.IDENTIFIER() != nil {
		label = ctx.IDENTIFIER().GetText()
	}
	return ast.NewGotoStmt(label)
}

func (visitor *GoParserVisitorImpl) VisitFallthroughStmt(ctx *antlr4.FallthroughStmtContext) interface{} {
	return ast.NewFallThroughStmt()
}

func (visitor *GoParserVisitorImpl) VisitDeferStmt(ctx *antlr4.DeferStmtContext) interface{} {
	expression := ctx.Expression().Accept(visitor).(*ast.Expression)
	return ast.NewDeferStmt(expression)
}

func (visitor *GoParserVisitorImpl) VisitIfStmt(ctx *antlr4.IfStmtContext) interface{} {
	var expression *ast.Expression = ctx.Expression().Accept(visitor).(*ast.Expression)
	var simpleStmt ast.SimpleStmt
	var block *ast.Block = ctx.Block(0).Accept(visitor).(*ast.Block)
	var elseif *ast.IfStmt
	var elseblock *ast.Block

	if ctx.SimpleStmt() != nil {
		simpleStmt = ctx.SimpleStmt().Accept(visitor).(ast.SimpleStmt)
	}
	if ctx.IfStmt() != nil {
		elseif = ctx.IfStmt().Accept(visitor).(*ast.IfStmt)
	}
	if len(ctx.AllBlock()) > 1 {
		elseblock = ctx.Block(1).Accept(visitor).(*ast.Block)
	}
	return ast.NewIfStmt(expression, simpleStmt, block, elseif, elseblock)
}

func (visitor *GoParserVisitorImpl) VisitSwitchStmt(ctx *antlr4.SwitchStmtContext) interface{} {
	if ctx.ExprSwitchStmt() != nil {
		return ctx.ExprSwitchStmt().Accept(visitor)
	}
	return ctx.TypeSwitchStmt().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitExprSwitchStmt(ctx *antlr4.ExprSwitchStmtContext) interface{} {
	var expression *ast.Expression
	var simpleStmt ast.SimpleStmt
	var exprCaseClauses []*ast.ExprCaseClause

	if ctx.Expression() != nil {
		expression = ctx.Expression().Accept(visitor).(*ast.Expression)
	}
	if ctx.SimpleStmt() != nil {
		simpleStmt = ctx.SimpleStmt().Accept(visitor).(ast.SimpleStmt)
	}
	for _, caseClause := range ctx.AllExprCaseClause() {
		exprCaseClauses = append(exprCaseClauses, caseClause.Accept(visitor).(*ast.ExprCaseClause))
	}
	return ast.NewExprSwitchStmt(expression, simpleStmt, exprCaseClauses)
}

func (visitor *GoParserVisitorImpl) VisitExprCaseClause(ctx *antlr4.ExprCaseClauseContext) interface{} {
	exprSwitchCase := ctx.ExprSwitchCase().Accept(visitor).(*ast.ExprSwitchCase)
	var statementList *ast.StatementList
	if ctx.StatementList() != nil {
		statementList = ctx.StatementList().Accept(visitor).(*ast.StatementList)
	}
	return ast.NewExprCaseClause(exprSwitchCase, statementList)
}

func (visitor *GoParserVisitorImpl) VisitExprSwitchCase(ctx *antlr4.ExprSwitchCaseContext) interface{} {
	var expressionList *ast.ExpressionList
	if ctx.ExpressionList() != nil {
		expressionList = ctx.ExpressionList().Accept(visitor).(*ast.ExpressionList)
	}
	return ast.NewExprSwitchCase(expressionList)
}

func (visitor *GoParserVisitorImpl) VisitTypeSwitchStmt(ctx *antlr4.TypeSwitchStmtContext) interface{} {
	var typeSwitchGuard *ast.TypeSwitchGuard = ctx.TypeSwitchGuard().Accept(visitor).(*ast.TypeSwitchGuard)
	var simpleStmt ast.SimpleStmt
	var typeCaseClauses []*ast.TypeCaseClause

	if ctx.SimpleStmt() != nil {
		simpleStmt = ctx.SimpleStmt().Accept(visitor).(ast.SimpleStmt)
	}
	for _, typecaseClause := range ctx.AllTypeCaseClause() {
		typeCaseClauses = append(typeCaseClauses, typecaseClause.Accept(visitor).(*ast.TypeCaseClause))
	}
	return ast.NewTypeSwitchStmt(typeSwitchGuard, simpleStmt, typeCaseClauses)
}

func (visitor *GoParserVisitorImpl) VisitTypeSwitchGuard(ctx *antlr4.TypeSwitchGuardContext) interface{} {
	var id string
	var primaryExpr *ast.PrimaryExpr = ctx.PrimaryExpr().Accept(visitor).(*ast.PrimaryExpr)
	if ctx.IDENTIFIER() != nil {
		id = ctx.IDENTIFIER().GetText()
	}
	return ast.NewTypeSwitchGuard(id, primaryExpr)
}

func (visitor *GoParserVisitorImpl) VisitTypeCaseClause(ctx *antlr4.TypeCaseClauseContext) interface{} {
	var typeSwitchCase *ast.TypeSwitchCase = ctx.TypeSwitchCase().Accept(visitor).(*ast.TypeSwitchCase)
	var statementList *ast.StatementList
	if ctx.StatementList() != nil {
		statementList = ctx.StatementList().Accept(visitor).(*ast.StatementList)
	}
	return ast.NewTypeCaseClause(typeSwitchCase, statementList)
}

func (visitor *GoParserVisitorImpl) VisitTypeSwitchCase(ctx *antlr4.TypeSwitchCaseContext) interface{} {
	var typeList *ast.TypeList
	if ctx.TypeList() != nil {
		typeList = ctx.TypeList().Accept(visitor).(*ast.TypeList)
	}
	return ast.NewTypeSwitchCase(typeList)
}

func (visitor *GoParserVisitorImpl) VisitTypeList(ctx *antlr4.TypeListContext) interface{} {
	var types []*ast.Type_
	for _, tree := range ctx.GetChildren() {
		switch child := tree.(type) {
		case *antlr4.Type_Context:
			types = append(types, child.Accept(visitor).(*ast.Type_))
		case antlr.TerminalNode:
			if child.GetText() == "nil" {
				types = append(types, nil)
			}
		}
	}
	return ast.NewTypeList(types)
}

func (visitor *GoParserVisitorImpl) VisitSelectStmt(ctx *antlr4.SelectStmtContext) interface{} {
	var commClauses []*ast.CommClause
	for _, clause := range ctx.AllCommClause() {
		commClauses = append(commClauses, clause.Accept(visitor).(*ast.CommClause))
	}
	return ast.NewSelectStmt(commClauses)
}

func (visitor *GoParserVisitorImpl) VisitCommClause(ctx *antlr4.CommClauseContext) interface{} {
	var commCase *ast.CommCase = ctx.CommCase().Accept(visitor).(*ast.CommCase)
	var statementList *ast.StatementList
	if ctx.StatementList() != nil {
		statementList = ctx.StatementList().Accept(visitor).(*ast.StatementList)
	}
	return ast.NewCommClause(commCase, statementList)
}

func (visitor *GoParserVisitorImpl) VisitCommCase(ctx *antlr4.CommCaseContext) interface{} {
	var sendStmt *ast.SendStmt
	var recvStmt *ast.RecvStmt

	if ctx.SendStmt() != nil {
		sendStmt = ctx.SendStmt().Accept(visitor).(*ast.SendStmt)
	}
	if ctx.RecvStmt() != nil {
		recvStmt = ctx.RecvStmt().Accept(visitor).(*ast.RecvStmt)
	}
	return ast.NewCommCase(sendStmt, recvStmt)
}

func (visitor *GoParserVisitorImpl) VisitRecvStmt(ctx *antlr4.RecvStmtContext) interface{} {
	var expressionList *ast.ExpressionList
	var identifierList *ast.IdentifierList
	var recvExpr *ast.Expression = ctx.GetRecvExpr().Accept(visitor).(*ast.Expression)

	if ctx.ExpressionList() != nil {
		expressionList = ctx.ExpressionList().Accept(visitor).(*ast.ExpressionList)
	}
	if ctx.IdentifierList() != nil {
		identifierList = ctx.IdentifierList().Accept(visitor).(*ast.IdentifierList)
	}
	return ast.NewRecvStmt(expressionList, identifierList, recvExpr)
}

func (visitor *GoParserVisitorImpl) VisitForStmt(ctx *antlr4.ForStmtContext) interface{} {
	var expression *ast.Expression
	var forClause *ast.ForClause
	var rangeClause *ast.RangeClause
	var block *ast.Block = ctx.Block().Accept(visitor).(*ast.Block)

	if ctx.Expression() != nil {
		expression = ctx.Expression().Accept(visitor).(*ast.Expression)
	}
	if ctx.ForClause() != nil {
		forClause = ctx.ForClause().Accept(visitor).(*ast.ForClause)
	}
	if ctx.RangeClause() != nil {
		rangeClause = ctx.RangeClause().Accept(visitor).(*ast.RangeClause)
	}
	return ast.NewForStmt(expression, forClause, rangeClause, block)
}

func (visitor *GoParserVisitorImpl) VisitForClause(ctx *antlr4.ForClauseContext) interface{} {
	var initStmt ast.SimpleStmt
	var expression *ast.Expression
	var postStmt ast.SimpleStmt

	if ctx.GetInitStmt() != nil {
		initStmt = ctx.GetInitStmt().Accept(visitor).(ast.SimpleStmt)
	}
	if ctx.Expression() != nil {
		expression = ctx.Expression().Accept(visitor).(*ast.Expression)
	}
	if ctx.GetPostStmt() != nil {
		postStmt = ctx.GetPostStmt().Accept(visitor).(ast.SimpleStmt)
	}
	return ast.NewForClause(initStmt, expression, postStmt)
}

func (visitor *GoParserVisitorImpl) VisitRangeClause(ctx *antlr4.RangeClauseContext) interface{} {
	var expressionList *ast.ExpressionList
	var identifierList *ast.IdentifierList
	var expression *ast.Expression = ctx.Expression().Accept(visitor).(*ast.Expression)

	if ctx.ExpressionList() != nil {
		expressionList = ctx.ExpressionList().Accept(visitor).(*ast.ExpressionList)
	}
	if ctx.IdentifierList() != nil {
		identifierList = ctx.IdentifierList().Accept(visitor).(*ast.IdentifierList)
	}
	return ast.NewRangeClause(expressionList, identifierList, expression)
}

func (visitor *GoParserVisitorImpl) VisitGoStmt(ctx *antlr4.GoStmtContext) interface{} {
	expression := ctx.Expression().Accept(visitor).(*ast.Expression)
	return ast.NewGoStmt(expression)
}

func (visitor *GoParserVisitorImpl) VisitType_(ctx *antlr4.Type_Context) interface{} {
	var typeName *ast.TypeName
	var typeLit ast.TypeLit
	var type_ *ast.Type_
	if ctx.TypeName() != nil {
		typeName = ctx.TypeName().Accept(visitor).(*ast.TypeName)
	}
	if ctx.TypeLit() != nil {
		typeLit = ctx.TypeLit().Accept(visitor).(ast.TypeLit)
	}
	if ctx.Type_() != nil {
		type_ = ctx.Type_().Accept(visitor).(*ast.Type_)
	}
	return ast.NewType_(typeName, typeLit, type_)
}

func (visitor *GoParserVisitorImpl) VisitTypeName(ctx *antlr4.TypeNameContext) interface{} {
	var qualifiedIdent *ast.QualifiedIdent
	var id string
	if ctx.QualifiedIdent() != nil {
		qualifiedIdent = ctx.QualifiedIdent().Accept(visitor).(*ast.QualifiedIdent)
	}
	if ctx.IDENTIFIER() != nil {
		id = ctx.IDENTIFIER().GetText()
	}
	return ast.NewTypeName(qualifiedIdent, id)
}

func (visitor *GoParserVisitorImpl) VisitTypeLit(ctx *antlr4.TypeLitContext) interface{} {
	if ctx.ArrayType() != nil {
		return ctx.ArrayType().Accept(visitor)
	}
	if ctx.StructType() != nil {
		return ctx.StructType().Accept(visitor)
	}
	if ctx.PointerType() != nil {
		return ctx.PointerType().Accept(visitor)
	}
	if ctx.FunctionType() != nil {
		return ctx.FunctionType().Accept(visitor)
	}
	if ctx.InterfaceType() != nil {
		return ctx.InterfaceType().Accept(visitor)
	}
	if ctx.SliceType() != nil {
		return ctx.SliceType().Accept(visitor)
	}
	if ctx.MapType() != nil {
		return ctx.MapType().Accept(visitor)
	}
	return ctx.ChannelType().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitArrayType(ctx *antlr4.ArrayTypeContext) interface{} {
	var arrayLength ast.ArrayLength = ctx.ArrayLength().Accept(visitor).(ast.ArrayLength)
	var elementType ast.ElementType = ctx.ElementType().Accept(visitor).(ast.ElementType)
	return ast.NewArrayType(arrayLength, elementType)
}

func (visitor *GoParserVisitorImpl) VisitArrayLength(ctx *antlr4.ArrayLengthContext) interface{} {
	return ctx.Expression().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitElementType(ctx *antlr4.ElementTypeContext) interface{} {
	return ctx.Type_().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitPointerType(ctx *antlr4.PointerTypeContext) interface{} {
	return ast.NewPointerType(ctx.Type_().Accept(visitor).(*ast.Type_))
}

func (visitor *GoParserVisitorImpl) VisitInterfaceType(ctx *antlr4.InterfaceTypeContext) interface{} {
	var methodSpecOrTypeNames []ast.IMethodSpecOrTypeName
	for _, tree := range ctx.GetChildren() {
		switch child := tree.(type) {
		case *antlr4.MethodSpecContext:
			methodSpecOrTypeNames = append(methodSpecOrTypeNames, child.Accept(visitor).(*ast.MethodSpec))
		case *antlr4.TypeNameContext:
			methodSpecOrTypeNames = append(methodSpecOrTypeNames, child.Accept(visitor).(*ast.TypeName))
		}
	}
	return ast.NewInterfaceType(methodSpecOrTypeNames)
}

func (visitor *GoParserVisitorImpl) VisitSliceType(ctx *antlr4.SliceTypeContext) interface{} {
	elementType := ctx.ElementType().Accept(visitor).(ast.ElementType)
	return ast.NewSliceType(elementType)
}

func (visitor *GoParserVisitorImpl) VisitMapType(ctx *antlr4.MapTypeContext) interface{} {
	type_ := ctx.Type_().Accept(visitor).(*ast.Type_)
	elementType := ctx.ElementType().Accept(visitor).(ast.ElementType)
	return ast.NewMapType(type_, elementType)
}

func (visitor *GoParserVisitorImpl) VisitChannelType(ctx *antlr4.ChannelTypeContext) interface{} {
	var title ast.ChannelTitle
	if ctx.GetCh() != nil {
		title = ast.CHAN
	} else if ctx.GetCh_re() != nil {
		title = ast.CHAN_RECEIVE
	} else if ctx.GetRe_ch() != nil {
		title = ast.RECEIVE_CHAN
	} else {
		panic("impossible.")
	}
	elementType := ctx.ElementType().Accept(visitor).(ast.ElementType)
	return ast.NewChannelType(title, elementType)
}

func (visitor *GoParserVisitorImpl) VisitMethodSpec(ctx *antlr4.MethodSpecContext) interface{} {
	var methodName string = ctx.IDENTIFIER().GetText()
	var parameters *ast.Parameters = ctx.Parameters().Accept(visitor).(*ast.Parameters)
	var result ast.Result
	if ctx.Result() != nil {
		result = ctx.Result().Accept(visitor).(ast.Result)
	}
	return ast.NewMethodSpec(methodName, parameters, result)
}

func (visitor *GoParserVisitorImpl) VisitFunctionType(ctx *antlr4.FunctionTypeContext) interface{} {
	var signature *ast.Signature = ctx.Signature().Accept(visitor).(*ast.Signature)
	return ast.NewFunctionType(signature)
}

func (visitor *GoParserVisitorImpl) VisitSignature(ctx *antlr4.SignatureContext) interface{} {
	var parameters *ast.Parameters = ctx.Parameters().Accept(visitor).(*ast.Parameters)
	var result ast.Result
	if ctx.Result() != nil {
		result = ctx.Result().Accept(visitor).(ast.Result)
	}
	return ast.NewSignature(parameters, result)
}

func (visitor *GoParserVisitorImpl) VisitResult(ctx *antlr4.ResultContext) interface{} {
	if ctx.Parameters() != nil {
		return ctx.Parameters().Accept(visitor)
	}
	return ctx.Type_().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitParameters(ctx *antlr4.ParametersContext) interface{} {
	var parameterDecls []*ast.ParameterDecl
	for _, parameter := range ctx.AllParameterDecl() {
		parameterDecls = append(parameterDecls, parameter.Accept(visitor).(*ast.ParameterDecl))
	}
	return ast.NewParameters(parameterDecls)
}

func (visitor *GoParserVisitorImpl) VisitParameterDecl(ctx *antlr4.ParameterDeclContext) interface{} {
	var identifierList *ast.IdentifierList
	var ellipsis bool = false
	var type_ *ast.Type_ = ctx.Type_().Accept(visitor).(*ast.Type_)
	if ctx.IdentifierList() != nil {
		identifierList = ctx.IdentifierList().Accept(visitor).(*ast.IdentifierList)
	}
	if ctx.ELLIPSIS() != nil {
		ellipsis = true
	}
	return ast.NewParameterDecl(identifierList, ellipsis, type_)
}

func (visitor *GoParserVisitorImpl) VisitExpression(ctx *antlr4.ExpressionContext) interface{} {
	var primaryExpr *ast.PrimaryExpr
	var expression *ast.Expression
	var expression2 *ast.Expression

	var unaryOp string
	var mulOp string
	var addOp string
	var relOp string
	var logicalAnd string
	var logicalOr string

	if ctx.PrimaryExpr() != nil {
		primaryExpr = ctx.PrimaryExpr().Accept(visitor).(*ast.PrimaryExpr)
	}
	if len(ctx.AllExpression()) > 0 {
		expression = ctx.Expression(0).Accept(visitor).(*ast.Expression)
	}
	if len(ctx.AllExpression()) > 1 {
		expression2 = ctx.Expression(1).Accept(visitor).(*ast.Expression)
	}
	if ctx.GetUnary_op() != nil {
		unaryOp = ctx.GetUnary_op().GetText()
	}
	if ctx.GetMul_op() != nil {
		mulOp = ctx.GetMul_op().GetText()
	}
	if ctx.GetAdd_op() != nil {
		addOp = ctx.GetAdd_op().GetText()
	}
	if ctx.GetRel_op() != nil {
		relOp = ctx.GetRel_op().GetText()
	}
	if ctx.LOGICAL_AND() != nil {
		logicalAnd = ctx.LOGICAL_AND().GetText()
	}
	if ctx.LOGICAL_OR() != nil {
		logicalOr = ctx.LOGICAL_OR().GetText()
	}
	return ast.NewExpressionBuilder().SetPrimaryExpr(primaryExpr).SetExpression(expression).SetExpression2(expression2).SetUnaryOp(unaryOp).SetMulOp(mulOp).SetAddOp(addOp).SetRelOp(relOp).SetLogicalOr(logicalOr).SetLogicalAnd(logicalAnd).Build()
}

func (visitor *GoParserVisitorImpl) VisitPrimaryExpr(ctx *antlr4.PrimaryExprContext) interface{} {
	var operand *ast.Operand
	var conversion *ast.Conversion
	var methodExpr *ast.MethodExpr
	var primaryExpr *ast.PrimaryExpr
	var id string
	var index *ast.Index
	var slice *ast.Slice
	var typeAssertion *ast.TypeAssertion
	var arguments *ast.Arguments

	if ctx.Operand() != nil {
		operand = ctx.Operand().Accept(visitor).(*ast.Operand)
	}
	if ctx.Conversion() != nil {
		conversion = ctx.Conversion().Accept(visitor).(*ast.Conversion)
	}
	if ctx.MethodExpr() != nil {
		methodExpr = ctx.MethodExpr().Accept(visitor).(*ast.MethodExpr)
	}
	if ctx.PrimaryExpr() != nil {
		primaryExpr = ctx.PrimaryExpr().Accept(visitor).(*ast.PrimaryExpr)
	}
	if ctx.IDENTIFIER() != nil {
		id = ctx.IDENTIFIER().GetText()
	}
	if ctx.Index() != nil {
		index = ctx.Index().Accept(visitor).(*ast.Index)
	}
	if ctx.Slice() != nil {
		slice = ctx.Slice().Accept(visitor).(*ast.Slice)
	}
	if ctx.TypeAssertion() != nil {
		typeAssertion = ctx.TypeAssertion().Accept(visitor).(*ast.TypeAssertion)
	}
	if ctx.Arguments() != nil {
		arguments = ctx.Arguments().Accept(visitor).(*ast.Arguments)
	}
	return ast.NewPrimaryExpr(operand, conversion, methodExpr, primaryExpr, id, index, slice, typeAssertion, arguments)
}

func (visitor *GoParserVisitorImpl) VisitConversion(ctx *antlr4.ConversionContext) interface{} {
	nonNamedType := ctx.NonNamedType().Accept(visitor).(*ast.NonNamedType)
	expression := ctx.Expression().Accept(visitor).(*ast.Expression)
	return ast.NewConversion(nonNamedType, expression)
}

func (visitor *GoParserVisitorImpl) VisitNonNamedType(ctx *antlr4.NonNamedTypeContext) interface{} {
	var typeLit ast.TypeLit
	var nonNamedType *ast.NonNamedType
	if ctx.TypeLit() != nil {
		typeLit = ctx.TypeLit().Accept(visitor).(ast.TypeLit)
	}
	if ctx.NonNamedType() != nil {
		nonNamedType = ctx.NonNamedType().Accept(visitor).(*ast.NonNamedType)
	}
	return ast.NewNonNamedType(typeLit, nonNamedType)
}

func (visitor *GoParserVisitorImpl) VisitOperand(ctx *antlr4.OperandContext) interface{} {
	var literal ast.Literal
	var operandName string
	var expression *ast.Expression
	if ctx.Literal() != nil {
		literal = ctx.Literal().Accept(visitor).(ast.Literal)
	}
	if ctx.OperandName() != nil {
		operandName = ctx.OperandName().GetText()
	}
	if ctx.Expression() != nil {
		expression = ctx.Expression().Accept(visitor).(*ast.Expression)
	}
	return ast.NewOperand(literal, operandName, expression)
}

func (visitor *GoParserVisitorImpl) VisitLiteral(ctx *antlr4.LiteralContext) interface{} {
	if ctx.BasicLit() != nil {
		return ctx.BasicLit().Accept(visitor)
	}
	if ctx.CompositeLit() != nil {
		return ctx.CompositeLit().Accept(visitor)
	}
	return ctx.FunctionLit().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitBasicLit(ctx *antlr4.BasicLitContext) interface{} {
	var isNil bool = ctx.NIL_LIT() != nil
	var integer *ast.Integer
	var string_ string
	var float string
	if ctx.Integer() != nil {
		integer = ctx.Integer().Accept(visitor).(*ast.Integer)
	}
	if ctx.String_() != nil {
		string_ = ctx.String_().GetText()
	}
	if ctx.FLOAT_LIT() != nil {
		float = ctx.FLOAT_LIT().GetText()
	}
	return ast.NewBasicLit(isNil, integer, string_, float)
}

func (visitor *GoParserVisitorImpl) VisitInteger(ctx *antlr4.IntegerContext) interface{} {
	if ctx.DECIMAL_LIT() != nil {
		return ast.NewInteger(ast.DECIMAL_LIT, ctx.DECIMAL_LIT().GetText())
	}
	if ctx.BINARY_LIT() != nil {
		return ast.NewInteger(ast.BINARY_LIT, ctx.BINARY_LIT().GetText())
	}
	if ctx.OCTAL_LIT() != nil {
		return ast.NewInteger(ast.OCTAL_LIT, ctx.OCTAL_LIT().GetText())
	}
	if ctx.HEX_LIT() != nil {
		return ast.NewInteger(ast.HEX_LIT, ctx.HEX_LIT().GetText())
	}
	if ctx.IMAGINARY_LIT() != nil {
		return ast.NewInteger(ast.IMAGINARY_LIT, ctx.IMAGINARY_LIT().GetText())
	}
	return ast.NewInteger(ast.DECIMAL_LIT, ctx.RUNE_LIT().GetText())
}

func (visitor *GoParserVisitorImpl) VisitOperandName(ctx *antlr4.OperandNameContext) interface{} {
	panic("It's impossible to run to here")
}

func (visitor *GoParserVisitorImpl) VisitQualifiedIdent(ctx *antlr4.QualifiedIdentContext) interface{} {
	id1 := ctx.IDENTIFIER(0).GetText()
	id2 := ctx.IDENTIFIER(1).GetText()
	return ast.NewQualifiedIdent(id1, id2)
}

func (visitor *GoParserVisitorImpl) VisitCompositeLit(ctx *antlr4.CompositeLitContext) interface{} {
	literalType := ctx.LiteralType().Accept(visitor).(*ast.LiteralType)
	literalValue := ctx.LiteralValue().Accept(visitor).(*ast.LiteralValue)
	return ast.NewCompositeLit(literalType, literalValue)
}

func (visitor *GoParserVisitorImpl) VisitLiteralType(ctx *antlr4.LiteralTypeContext) interface{} {
	var structType *ast.StructType
	var arrayType *ast.ArrayType
	var elementType ast.ElementType
	var sliceType *ast.SliceType
	var mapType *ast.MapType
	var typeName *ast.TypeName

	if ctx.StructType() != nil {
		structType = ctx.StructType().Accept(visitor).(*ast.StructType)
	}
	if ctx.ArrayType() != nil {
		arrayType = ctx.ArrayType().Accept(visitor).(*ast.ArrayType)
	}
	if ctx.ElementType() != nil {
		elementType = ctx.ElementType().Accept(visitor).(ast.ElementType)
	}
	if ctx.SliceType() != nil {
		sliceType = ctx.SliceType().Accept(visitor).(*ast.SliceType)
	}
	if ctx.MapType() != nil {
		mapType = ctx.MapType().Accept(visitor).(*ast.MapType)
	}
	if ctx.TypeName() != nil {
		typeName = ctx.TypeName().Accept(visitor).(*ast.TypeName)
	}
	return ast.NewLiteralType(structType, arrayType, elementType, sliceType, mapType, typeName)
}

func (visitor *GoParserVisitorImpl) VisitLiteralValue(ctx *antlr4.LiteralValueContext) interface{} {
	var elementList *ast.ElementList
	if ctx.ElementList() != nil {
		elementList = ctx.ElementList().Accept(visitor).(*ast.ElementList)
	}
	return ast.NewLiteralValue(elementList)
}

func (visitor *GoParserVisitorImpl) VisitElementList(ctx *antlr4.ElementListContext) interface{} {
	var keyedElements []*ast.KeyedElement
	for _, element := range ctx.AllKeyedElement() {
		keyedElements = append(keyedElements, element.Accept(visitor).(*ast.KeyedElement))
	}
	return ast.NewElementList(keyedElements)
}

func (visitor *GoParserVisitorImpl) VisitKeyedElement(ctx *antlr4.KeyedElementContext) interface{} {
	var key ast.Key
	element := ctx.Element().Accept(visitor).(ast.Element)
	if ctx.Key() != nil {
		key = ctx.Key().Accept(visitor).(ast.Key)
	}
	return ast.NewKeyedElement(key, element)
}

func (visitor *GoParserVisitorImpl) VisitKey(ctx *antlr4.KeyContext) interface{} {
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(visitor)
	}
	return ctx.LiteralValue().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitElement(ctx *antlr4.ElementContext) interface{} {
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(visitor)
	}
	return ctx.LiteralValue().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitStructType(ctx *antlr4.StructTypeContext) interface{} {
	var fieldDecls []*ast.FieldDecl
	for _, decl := range ctx.AllFieldDecl() {
		fieldDecls = append(fieldDecls, decl.Accept(visitor).(*ast.FieldDecl))
	}
	return ast.NewStructType(fieldDecls)
}

func (visitor *GoParserVisitorImpl) VisitFieldDecl(ctx *antlr4.FieldDeclContext) interface{} {
	var annotation string
	//if ctx.ANNOTATION() != nil {
	//	annotation = ctx.ANNOTATION().GetText()
	//}
	var identifierList *ast.IdentifierList
	var type_ *ast.Type_
	var embeddedField *ast.EmbeddedField
	var tag string

	if ctx.IdentifierList() != nil {
		identifierList = ctx.IdentifierList().Accept(visitor).(*ast.IdentifierList)
	}
	if ctx.Type_() != nil {
		type_ = ctx.Type_().Accept(visitor).(*ast.Type_)
	}
	if ctx.EmbeddedField() != nil {
		embeddedField = ctx.EmbeddedField().Accept(visitor).(*ast.EmbeddedField)
	}
	if ctx.GetTag() != nil {
		tag = ctx.GetTag().GetText()
	}
	return ast.NewFieldDecl(annotation, identifierList, type_, embeddedField, tag)
}

func (visitor *GoParserVisitorImpl) VisitString_(ctx *antlr4.String_Context) interface{} {
	panic("It is impossible to run to this line")
}

func (visitor *GoParserVisitorImpl) VisitEmbeddedField(ctx *antlr4.EmbeddedFieldContext) interface{} {
	star := ctx.STAR() != nil
	typeName := ctx.TypeName().Accept(visitor).(*ast.TypeName)
	return ast.NewEmbeddedField(star, typeName)
}

func (visitor *GoParserVisitorImpl) VisitFunctionLit(ctx *antlr4.FunctionLitContext) interface{} {
	signature := ctx.Signature().Accept(visitor).(*ast.Signature)
	block := ctx.Block().Accept(visitor).(*ast.Block)
	return ast.NewFunctionLit(signature, block)
}

func (visitor *GoParserVisitorImpl) VisitIndex(ctx *antlr4.IndexContext) interface{} {
	expression := ctx.Expression().Accept(visitor).(*ast.Expression)
	return ast.NewIndex(expression)
}

func (visitor *GoParserVisitorImpl) VisitSlice(ctx *antlr4.SliceContext) interface{} {
	var expression1 *ast.Expression
	var expression2 *ast.Expression
	var expression3 *ast.Expression

	if len(ctx.AllExpression()) > 0 {
		expression1 = ctx.Expression(0).Accept(visitor).(*ast.Expression)
	}
	if len(ctx.AllExpression()) > 1 {
		expression2 = ctx.Expression(1).Accept(visitor).(*ast.Expression)
	}
	if len(ctx.AllExpression()) > 2 {
		expression3 = ctx.Expression(2).Accept(visitor).(*ast.Expression)
	}
	return ast.NewSlice(expression1, expression2, expression3)
}

func (visitor *GoParserVisitorImpl) VisitTypeAssertion(ctx *antlr4.TypeAssertionContext) interface{} {
	return ast.NewTypeAssertion(ctx.Type_().Accept(visitor).(*ast.Type_))
}

func (visitor *GoParserVisitorImpl) VisitArguments(ctx *antlr4.ArgumentsContext) interface{} {
	var expressionList *ast.ExpressionList
	var nonNamedType *ast.NonNamedType
	var ellipsis bool = ctx.ELLIPSIS() != nil
	if ctx.ExpressionList() != nil {
		expressionList = ctx.ExpressionList().Accept(visitor).(*ast.ExpressionList)
	}
	if ctx.NonNamedType() != nil {
		nonNamedType = ctx.NonNamedType().Accept(visitor).(*ast.NonNamedType)
	}
	return ast.NewArguments(expressionList, nonNamedType, ellipsis)
}

func (visitor *GoParserVisitorImpl) VisitMethodExpr(ctx *antlr4.MethodExprContext) interface{} {
	nonNamedType := ctx.NonNamedType().Accept(visitor).(*ast.NonNamedType)
	id := ctx.IDENTIFIER().GetText()
	return ast.NewMethodExpr(nonNamedType, id)
}

func (visitor *GoParserVisitorImpl) VisitReceiverType(ctx *antlr4.ReceiverTypeContext) interface{} {
	return ctx.Type_().Accept(visitor)
}

func (visitor *GoParserVisitorImpl) VisitEos(ctx *antlr4.EosContext) interface{} {
	println("++++++++", ctx.GetText())
	return ast.NewEos(ctx.GetText())
}

var _ antlr4.GoParserVisitor = (*GoParserVisitorImpl)(nil)
