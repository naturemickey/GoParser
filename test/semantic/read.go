package semantic

import (
	"GoParser/parser/ast"
	"reflect"
	"strings"
)

var o struct{}

type Walker struct {
	goset map[*ast.Operand]struct{}
}

func (w *Walker) WalkSourceFile(tree *ast.SourceFile) {
	if w.goset != nil {
		panic("你用错了")
	} else {
		w.goset = map[*ast.Operand]struct{}{}
	}
	scope := NewScopeRoot()
	for _, importDecl := range tree.ImportDecls() {
		for _, importSpec := range importDecl.ImportSpecs() {
			name := importSpec.Alias()
			if name != "" {
				tree := importSpec
				type_ := PACKAGE
				typeLiteral := importSpec.ImportPath().String()
				if typeLiteral == "\"context\"" {
					scope.AddName(name, NewName(name, tree, type_, typeLiteral))
				}
			}
		}
	}
	for _, fmd := range tree.FunctionOrMethodOrDeclarations() {
		switch child := fmd.(type) {
		case *ast.TypeDecl:
			w.WalkTypeDecl(scope, child, false)
		}
	}
	for _, fmd := range tree.FunctionOrMethodOrDeclarations() {
		switch child := fmd.(type) {
		case *ast.FunctionDecl:
			name := child.FunName()
			tree := child
			type_ := FUNCATION
			typeLiteral := ""
			scope.AddName(name, NewName(name, tree, type_, typeLiteral))
		case *ast.MethodDecl:
			name := child.FunName()
			tree := child
			type_ := METHOD
			typeLiteral := ""
			scope.AddName(name, NewName(name, tree, type_, typeLiteral))
		//case ast.Declaration:
		case *ast.ConstDecl:
			w.WalkConstDecl(scope, child, false)
		case *ast.VarDecl:
			w.WalkVarDecl(scope, child, false)
		}
	}
	scope.forNames(func(id string, name *Name) {
		if name.Type_() == FUNCATION {
			w.WalkFunctionDecl(scope, name.tree.(*ast.FunctionDecl), false)
		} else if name.Type_() == METHOD {
			w.WalkMethodDecl(scope, name.tree.(*ast.MethodDecl), false)
		}
	})
}

func (w *Walker) WalkTypeDecl(scope *Scope, typeDecl *ast.TypeDecl, isGo bool) {
	for _, typeSpec := range typeDecl.TypeSpecs() {
		name := typeSpec.Id()
		tree := typeSpec
		type_ := TYPE
		typeLiteral := typeSpec.Type_().String()
		scope.AddName(name, NewName(name, tree, type_, typeLiteral))
	}
}

func (w *Walker) WalkConstDecl(scope *Scope, constDecl *ast.ConstDecl, isGo bool) {
	type_ := VARIABLE
	for _, constSpec := range constDecl.ConstSpecs() {
		var typeLiteral string
		if constSpec.Type_() != nil {
			typeLiteral = constSpec.Type_().String()
		}
		for i, name := range constSpec.IdentifierList().Identifiers() {
			if constSpec.ExpressionList() != nil && i < len(constSpec.ExpressionList().Expressions()) { // 这个判断也是一个bug
				tree := constSpec.ExpressionList().Expressions()[i]
				scope.AddName(name, NewName(name, tree, type_, typeLiteral))
			}
		}
	}
}

func (w *Walker) WalkVarDecl(scope *Scope, varDecl *ast.VarDecl, isGo bool) {
	type_ := VARIABLE
	for _, constSpec := range varDecl.VarSpecs() {
		var typeLiteral string
		if constSpec.Type_() != nil {
			typeLiteral = constSpec.Type_().String()
		}
		for i, name := range constSpec.IdentifierList().Identifiers() {
			if constSpec.ExpressionList() != nil && i < len(constSpec.ExpressionList().Expressions()) { // 这个判断也是一个bug，先放过
				tree := constSpec.ExpressionList().Expressions()[i]
				scope.AddName(name, NewName(name, tree, type_, typeLiteral))
			}
		}
	}
}

func (w *Walker) WalkFunctionDecl(scope *Scope, functionDecl *ast.FunctionDecl, isGo bool) {
	scope = scope.NewChildScope(false)

	w.WalkSignature(scope, functionDecl.Signature(), isGo)
	w.WalkBlock(scope, functionDecl.Block(), false, isGo)
}

func (w *Walker) WalkMethodDecl(scope *Scope, methodDecl *ast.MethodDecl, isGo bool) {
	scope = scope.NewChildScope(false)

	w.WalkParameters(scope, methodDecl.Receiver().(*ast.Parameters), isGo)
	w.WalkSignature(scope, methodDecl.Signature(), isGo)
	w.WalkBlock(scope, methodDecl.Block(), false, isGo)
}

func (w *Walker) WalkSignature(scope *Scope, signature *ast.Signature, isGo bool) {
	w.WalkParameters(scope, signature.Parameters(), isGo)
}

func (w *Walker) WalkParameters(scope *Scope, parameters *ast.Parameters, isGo bool) {
	if parameters == nil {
		return
	}
	type_ := VARIABLE
	var tree ast.INode = nil
	for _, parameterDecl := range parameters.ParameterDecls() {
		if parameterDecl != nil {
			typeLiteral := parameterDecl.Type_().String()
			if parameterDecl.IdentifierList() != nil {
				for _, name := range parameterDecl.IdentifierList().Identifiers() {
					scope.AddName(name, NewName(name, tree, type_, typeLiteral))
				}
			}
		}
	}
}
func (w *Walker) WalkBlock(scope *Scope, block *ast.Block, newScope bool, isGo bool) {
	if block == nil {
		return
	}
	if newScope {
		scope = scope.NewChildScope(false)
	}
	if block.StatementList() != nil {
		for _, statement := range block.StatementList().Statements() {
			w.WalkStatement(scope, statement, isGo)
		}
	}
}

func (w *Walker) WalkStatement(scope *Scope, statement ast.Statement, isGo bool) {
	if statement == nil || reflect.ValueOf(statement).IsNil() {
		return
	}
	switch stmt := statement.(type) {
	case *ast.ConstDecl:
		w.WalkConstDecl(scope, stmt, isGo)
	case *ast.TypeDecl:
		w.WalkTypeDecl(scope, stmt, isGo)
	case *ast.VarDecl:
		w.WalkVarDecl(scope, stmt, isGo)
	case *ast.LabeledStmt:
		w.WalkStatement(scope, stmt.Statement(), isGo)
	case *ast.SendStmt:
		w.WalkExpression(scope, stmt.Channel(), isGo)
		w.WalkExpression(scope, stmt.Expression(), isGo)
	case *ast.IncDecStmt:
		w.WalkExpression(scope, stmt.Expression(), isGo)
	case *ast.Assignment:
		w.WalkExpressionList(scope, stmt.ExpressionListLeft(), isGo)
		w.WalkExpressionList(scope, stmt.ExpressionListRight(), isGo)
	case *ast.Expression:
		w.WalkExpression(scope, stmt, isGo)
	case *ast.ShortValDecl:
		w.WalkShortValDecl(scope, stmt, isGo)
	case *ast.GoStmt:
		w.WalkExpression(scope.NewChildScope(true), stmt.Expression(), true)
	case *ast.ReturnStmt:
		w.WalkExpressionList(scope, stmt.ExpressionList(), isGo)
	case *ast.BreakStmt: // don't need to do anything
	case *ast.ContinueStmt: // don't need to do anything
	case *ast.GotoStmt: // don't need to do anything
	case *ast.FallThroughStmt: // don't need to do anything
	case *ast.Block:
		w.WalkBlock(scope, stmt, true, isGo)
	case *ast.IfStmt:
		w.WalkIfStmt(scope, stmt, isGo)
	case *ast.ExprSwitchStmt:
		newScope := scope.NewChildScope(false)
		w.WalkStatement(newScope, stmt.SimpleStmt(), isGo)
		w.WalkExpression(newScope, stmt.Expression(), isGo)
		for _, caseClause := range stmt.ExprCaseClauses() {
			w.WalkExprCaseClause(scope, caseClause, isGo)
		}
	case *ast.TypeSwitchStmt:
		newScope := scope.NewChildScope(false)
		w.WalkStatement(newScope, stmt.SimpleStmt(), isGo)
		w.WalkTypeSwitchGuard(newScope, stmt.TypeSwitchGuard(), isGo)
		for _, typeCaseClause := range stmt.TypeCaseClauses() {
			w.WalkTypeCaseClause(scope, typeCaseClause, isGo)
		}
	case *ast.SelectStmt:
		for _, clause := range stmt.CommClauses() {
			w.WalkCommClause(scope, clause, isGo)
		}
	case *ast.ForStmt:
		newScope := scope.NewChildScope(false)
		w.WalkExpression(newScope, stmt.Expression(), isGo)
		w.WalkForClause(newScope, stmt.ForClause(), isGo)
		w.WalkRangeClause(newScope, stmt.RangeClause(), isGo)
		w.WalkBlock(newScope, stmt.Block(), false, isGo)
	case *ast.DeferStmt:
		w.WalkExpression(scope, stmt.Expression(), isGo)
	default:
		panic(statement.String())
	}
}

func (w *Walker) WalkRangeClause(scope *Scope, rangeClause *ast.RangeClause, isGo bool) {
	if rangeClause == nil {
		return
	}
	w.WalkExpressionList(scope, rangeClause.ExpressionList(), isGo)
	if rangeClause.IdentifierList() != nil {
		for _, id := range rangeClause.IdentifierList().Identifiers() {
			scope.AddName(id, NewName(id, nil, VARIABLE, ""))
		}
	}
	w.WalkExpression(scope, rangeClause.Expression(), isGo)
}

func (w *Walker) WalkForClause(scope *Scope, forClause *ast.ForClause, isGo bool) {
	if forClause == nil {
		return
	}
	w.WalkStatement(scope, forClause.InitStmt(), isGo)
	w.WalkExpression(scope, forClause.Expression(), isGo)
	w.WalkStatement(scope, forClause.PostStmt(), isGo)
}

func (w *Walker) WalkCommClause(scope *Scope, commClause *ast.CommClause, isGo bool) {
	if commClause == nil {
		return
	}
	scope = scope.NewChildScope(false)
	w.WalkCommCase(scope, commClause.CommCase(), isGo)
	w.WalkStatementList(scope, commClause.StatementList(), isGo)
}

func (w *Walker) WalkStatementList(scope *Scope, statementList *ast.StatementList, isGo bool) {
	if statementList == nil {
		return
	}
	for _, statement := range statementList.Statements() {
		w.WalkStatement(scope, statement, isGo)
	}
}

func (w *Walker) WalkCommCase(scope *Scope, commCase *ast.CommCase, isGo bool) {
	if commCase == nil {
		return
	}
	w.WalkSendStmt(scope, commCase.SendStmt(), isGo)
	w.WalkRecvStmt(scope, commCase.RecvStmt(), isGo)
}

func (w *Walker) WalkSendStmt(scope *Scope, stmt *ast.SendStmt, isGo bool) {
	if stmt == nil {
		return
	}
	w.WalkExpression(scope, stmt.Channel(), isGo)
	w.WalkExpression(scope, stmt.Expression(), isGo)
}

func (w *Walker) WalkRecvStmt(scope *Scope, stmt *ast.RecvStmt, isGo bool) {
	if stmt == nil {
		return
	}
	w.WalkExpressionList(scope, stmt.ExpressionList(), isGo)
	if stmt.IdentifierList() != nil {
		for _, id := range stmt.IdentifierList().Identifiers() {
			scope.AddName(id, NewName(id, nil, VARIABLE, ""))
		}
	}
	w.WalkExpression(scope, stmt.RecvExpr(), isGo)
}

func (w *Walker) WalkExprCaseClause(scope *Scope, exprCaseClause *ast.ExprCaseClause, isGo bool) {
	if exprCaseClause == nil {
		return
	}
	scope = scope.NewChildScope(false)
	w.WalkExprSwitchCase(scope, exprCaseClause.ExprSwitchCase(), isGo)
	w.WalkStatementList(scope, exprCaseClause.StatementList(), isGo)
}

func (w *Walker) WalkExprSwitchCase(scope *Scope, switchCase *ast.ExprSwitchCase, isGo bool) {
	if switchCase == nil {
		return
	}
	w.WalkExpressionList(scope, switchCase.ExpressionList(), isGo)
}

func (w *Walker) WalkTypeCaseClause(scope *Scope, exprCaseClause *ast.TypeCaseClause, isGo bool) {
	if exprCaseClause == nil {
		return
	}
	scope = scope.NewChildScope(false)
	w.WalkTypeSwitchCase(scope, exprCaseClause.TypeSwitchCase(), isGo)
	w.WalkStatementList(scope, exprCaseClause.StatementList(), isGo)
}

func (w *Walker) WalkTypeSwitchCase(scope *Scope, switchCase *ast.TypeSwitchCase, isGo bool) {
	if switchCase == nil {
		return
	}
	w.WalkTypeList(scope, switchCase.TypeList(), isGo)
}

func (w *Walker) WalkTypeList(scope *Scope, typeList *ast.TypeList, isGo bool) {
	if typeList == nil {
		return
	}
	for _, type_ := range typeList.Types() {
		w.WalkType_(scope, type_, isGo)
	}
}

func (w *Walker) WalkType_(scope *Scope, type_ *ast.Type_, isGo bool) {
	if type_ == nil {
		return
	}
	// TODO 好像并不需要，走走再看
}

func (w *Walker) WalkTypeName(scope *Scope, name *ast.TypeName, isGo bool) {

}

func (w *Walker) WalkTypeSwitchGuard(scope *Scope, typeSwitchGuard *ast.TypeSwitchGuard, isGo bool) {
	if typeSwitchGuard == nil {
		return
	}
	w.WalkPrimaryExpr(scope, typeSwitchGuard.PrimaryExpr(), isGo)
}

func (w *Walker) WalkPrimaryExpr(scope *Scope, expr *ast.PrimaryExpr, isGo bool) {
	if expr == nil {
		return
	}
	w.WalkOperand(scope, expr.Operand(), isGo)
	w.WalkConversion(scope, expr.Conversion(), isGo)
	w.WalkMethodExpr(scope, expr.MethodExpr(), isGo)
	w.WalkPrimaryExpr(scope, expr.PrimaryExpr(), isGo)
	w.WalkIndex(scope, expr.Index(), isGo)
	w.WalkSlice(scope, expr.Slice(), isGo)
	w.WalkTypeAssertion(scope, expr.TypeAssertion(), isGo)
	w.WalkArguments(scope, expr.Arguments(), isGo)
}

func (w *Walker) WalkOperand(scope *Scope, stmt *ast.Operand, isGo bool) {
	if stmt == nil {
		return
	}
	w.WalkLiteral(scope, stmt.Literal(), isGo)

	if isGo {
		_, exist := w.goset[stmt]
		if exist {
			return
		} else {
			w.goset[stmt] = o
		}

		name := stmt.OperandName()
		n, goLevel := scope.GetNameByName(name)
		if n != nil {
			if goLevel != scope.goLevel {
				if n.TypeLiteral() == "context.Context" {
					panic("在协程里用了外部的context.Context")
				} else if n.TypeLiteral() == "Context" {
					contexAlias, _ := scope.GetNameByName(".")
					if contexAlias != nil && contexAlias.type_ == PACKAGE && contexAlias.typeLiteral == "\"context\"" {
						panic("在协程里用了外部的context.Context")
					}
				} else {
					pkgn, _ := scope.GetNameByName(n.TypeLiteral())
					if pkgn != nil && pkgn.typeLiteral == "context.Context" {
						panic("在协程里用了外部的context.Context")
					}
				}
			}
			if n.tree != nil {
				if goLevel == scope.goLevel || n.type_ != VARIABLE || strings.HasPrefix(n.typeLiteral, "func(") {
					switch node := n.tree.(type) {
					case *ast.ExpressionList:
						w.WalkExpressionList(scope, node, isGo)
					case *ast.Expression:
						w.WalkExpression(scope, node, isGo)
					case *ast.FunctionDecl:
						w.WalkFunctionDecl(scope, node, isGo)
					}
				}
			}
		}
	}
	w.WalkExpression(scope, stmt.Expression(), isGo)
}

func (w *Walker) WalkLiteral(scope *Scope, stmt ast.Literal, isGo bool) {
	if stmt == nil {
		return
	}
	switch tree := stmt.(type) {
	case *ast.BasicLit:
		w.WalkBasicLit(scope, tree, isGo)
	case *ast.CompositeLit:
		w.WalkCompositeLit(scope, tree, isGo)
	case *ast.FunctionLit:
		w.WalkFunctionLit(scope, tree, isGo)
	}
}

func (w *Walker) WalkBasicLit(scope *Scope, stmt *ast.BasicLit, isGo bool) {
	// 看起来不用搞这个函数
}

func (w *Walker) WalkFunctionLit(scope *Scope, stmt *ast.FunctionLit, isGo bool) {
	if stmt == nil {
		return
	}
	scope = scope.NewChildScope(false)
	w.WalkSignature(scope, stmt.Signature(), isGo)
	w.WalkBlock(scope, stmt.Block(), false, isGo)
}

func (w *Walker) WalkCompositeLit(scope *Scope, stmt *ast.CompositeLit, isGo bool) {
	if stmt == nil {
		return
	}
	w.WalkLiteralType(scope, stmt.LiteralType(), isGo)
	w.WalkLiteralValue(scope, stmt.LiteralValue(), isGo)
}

func (w *Walker) WalkLiteralType(scope *Scope, stmt *ast.LiteralType, isGo bool) {
	// 看起来不用搞这个函数
}

func (w *Walker) WalkLiteralValue(scope *Scope, stmt *ast.LiteralValue, isGo bool) {
	if stmt == nil {
		return
	}
	w.WalkElementList(scope, stmt.ElementList(), isGo)
}

func (w *Walker) WalkElementList(scope *Scope, stmt *ast.ElementList, isGo bool) {
	if stmt == nil {
		return
	}
	for _, element := range stmt.KeyedElements() {
		w.WalkKeyedElement(scope, element, isGo)
	}
}

func (w *Walker) WalkKeyedElement(scope *Scope, stmt *ast.KeyedElement, isGo bool) {
	if stmt == nil {
		return
	}
	if stmt.Key() != nil && !reflect.ValueOf(stmt.Key()).IsNil() {
		switch key := stmt.Key().(type) {
		case *ast.Expression:
			w.WalkExpression(scope, key, isGo)
		case *ast.LiteralValue:
			w.WalkLiteralValue(scope, key, isGo)
		default:
			panic("impossible")
		}
	}
	switch element := stmt.Element().(type) {
	case *ast.Expression:
		w.WalkExpression(scope, element, isGo)
	case *ast.LiteralValue:
		w.WalkLiteralValue(scope, element, isGo)
	}
}

func (w *Walker) WalkConversion(scope *Scope, stmt *ast.Conversion, isGo bool) {
	if stmt == nil {
		return
	}
	w.WalkNonNamedType(scope, stmt.NonNamedType(), isGo)
	w.WalkExpression(scope, stmt.Expression(), isGo)
}

func (w *Walker) WalkNonNamedType(scope *Scope, stmt *ast.NonNamedType, isGo bool) {
	if stmt == nil {
		return
	}
	w.WalkTypeLit(scope, stmt.TypeLit(), isGo)
	w.WalkNonNamedType(scope, stmt.NonNamedType(), isGo)
}

func (w *Walker) WalkTypeLit(scope *Scope, stmt ast.TypeLit, isGo bool) {
	// todo
}

func (w *Walker) WalkMethodExpr(scope *Scope, stmt *ast.MethodExpr, isGo bool) {

}

func (w *Walker) WalkIndex(scope *Scope, stmt *ast.Index, isGo bool) {

}

func (w *Walker) WalkSlice(scope *Scope, stmt *ast.Slice, isGo bool) {

}

func (w *Walker) WalkTypeAssertion(scope *Scope, stmt *ast.TypeAssertion, isGo bool) {

}

func (w *Walker) WalkArguments(scope *Scope, stmt *ast.Arguments, isGo bool) {
	//if scope.parent != nil {
	//	goLevel := scope.goLevel
	//	parentGoLevel := scope.parent.goLevel
	//	if goLevel != parentGoLevel {
	//		return
	//	}
	//}

	if stmt == nil {
		return
	}
	w.WalkExpressionList(scope, stmt.ExpressionList(), isGo)
	w.WalkNonNamedType(scope, stmt.NonNamedType(), isGo)
}

func (w *Walker) WalkIfStmt(scope *Scope, stmt *ast.IfStmt, isGo bool) {
	if stmt == nil {
		return
	}
	newScope := scope.NewChildScope(false)
	w.WalkStatement(newScope, stmt.SimpleStmt(), isGo)
	w.WalkExpression(newScope, stmt.Expression(), isGo)
	w.WalkBlock(newScope, stmt.Block(), true, isGo)
	w.WalkIfStmt(newScope, stmt.Elseif(), isGo)
	w.WalkBlock(newScope, stmt.Elseblock(), true, isGo)
}

func (w *Walker) WalkShortValDecl(scope *Scope, shortValDecl *ast.ShortValDecl, isGo bool) {
	if shortValDecl == nil {
		return
	}
	for i, name := range shortValDecl.IdentifierList().Identifiers() {
		if i < len(shortValDecl.ExpressionList().Expressions()) { // 这个if是一个bug，没事，先这样干
			tree := shortValDecl.ExpressionList().Expressions()[i]
			type_ := VARIABLE
			typeLiteral := tree.String()
			scope.AddName(name, NewName(name, tree, type_, typeLiteral))

			w.WalkExpression(scope, tree, isGo)
		}
	}
}

func (w *Walker) WalkExpressionList(scope *Scope, expressionList *ast.ExpressionList, isGo bool) {
	if expressionList == nil {
		return
	}
	for _, expression := range expressionList.Expressions() {
		w.WalkExpression(scope, expression, isGo)
	}
}

func (w *Walker) WalkExpression(scope *Scope, expression *ast.Expression, isGo bool) {
	if expression == nil {
		return
	}
	w.WalkPrimaryExpr(scope, expression.PrimaryExpr(), isGo)
	w.WalkExpression(scope, expression.Expression(), isGo)
	w.WalkExpression(scope, expression.Expression2(), isGo)
}
