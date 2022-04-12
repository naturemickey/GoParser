package semantic

import (
	"GoParser/parser/ast"
	"reflect"
)

func WalkSourceFile(tree ast.SourceFile) {
	scope := NewScopeRoot()
	for _, fmd := range tree.FunctionOrMethodOrDeclarations() {
		switch child := fmd.(type) {
		case *ast.TypeDecl:
			WalkTypeDecl(scope, child, false)
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
			WalkConstDecl(scope, child, false)
		case *ast.VarDecl:
			WalkVarDecl(scope, child, false)
		}
	}
	scope.forNames(func(id string, name *Name) {
		if name.Type_() == FUNCATION {
			WalkFunctionDecl(scope, name.tree.(*ast.FunctionDecl), false)
		} else if name.Type_() == METHOD {
			WalkMethodDecl(scope, name.tree.(*ast.MethodDecl), false)
		}
	})
}

func WalkTypeDecl(scope *Scope, typeDecl *ast.TypeDecl, isGo bool) {
	for _, typeSpec := range typeDecl.TypeSpecs() {
		name := typeSpec.Id()
		tree := typeSpec
		type_ := TYPE
		typeLiteral := typeSpec.String()
		scope.AddName(name, NewName(name, tree, type_, typeLiteral))
	}
}

func WalkConstDecl(scope *Scope, constDecl *ast.ConstDecl, isGo bool) {
	type_ := VARIABLE
	for _, constSpec := range constDecl.ConstSpecs() {
		var typeLiteral string
		if constSpec.Type_() != nil {
			typeLiteral = constSpec.Type_().String()
		}
		for i, name := range constSpec.IdentifierList().Identifiers() {
			tree := constSpec.ExpressionList().Expressions()[i]
			scope.AddName(name, NewName(name, tree, type_, typeLiteral))
		}
	}
}

func WalkVarDecl(scope *Scope, varDecl *ast.VarDecl, isGo bool) {
	type_ := VARIABLE
	for _, constSpec := range varDecl.VarSpecs() {
		var typeLiteral string
		if constSpec.Type_() != nil {
			typeLiteral = constSpec.Type_().String()
		}
		for i, name := range constSpec.IdentifierList().Identifiers() {
			tree := constSpec.ExpressionList().Expressions()[i]
			scope.AddName(name, NewName(name, tree, type_, typeLiteral))
		}
	}
}

func WalkFunctionDecl(scope *Scope, functionDecl *ast.FunctionDecl, isGo bool) {
	scope = scope.NewChildScope()

	WalkSignature(scope, functionDecl.Signature(), isGo)
	WalkBlock(scope, functionDecl.Block(), false, isGo)
}

func WalkMethodDecl(scope *Scope, methodDecl *ast.MethodDecl, isGo bool) {
	scope = scope.NewChildScope()

	WalkParameters(scope, methodDecl.Receiver().(*ast.Parameters), isGo)
	WalkSignature(scope, methodDecl.Signature(), isGo)
	WalkBlock(scope, methodDecl.Block(), false, isGo)
}

func WalkSignature(scope *Scope, signature *ast.Signature, isGo bool) {
	WalkParameters(scope, signature.Parameters(), isGo)
}

func WalkParameters(scope *Scope, parameters *ast.Parameters, isGo bool) {
	type_ := VARIABLE
	var tree ast.INode = nil
	for _, parameterDecl := range parameters.ParameterDecls() {
		typeLiteral := parameterDecl.Type_().String()
		for _, name := range parameterDecl.IdentifierList().Identifiers() {
			scope.AddName(name, NewName(name, tree, type_, typeLiteral))
		}
	}
}
func WalkBlock(scope *Scope, block *ast.Block, newScope bool, isGo bool) {
	if block == nil {
		return
	}
	if newScope {
		scope = scope.NewChildScope()
	}
	if block.StatementList() != nil {
		for _, statement := range block.StatementList().Statements() {
			WalkStatement(scope, statement, isGo)
		}
	}
}

func WalkStatement(scope *Scope, statement ast.Statement, isGo bool) {
	if statement == nil || reflect.ValueOf(statement).IsNil() {
		return
	}
	switch stmt := statement.(type) {
	case *ast.ConstDecl:
		WalkConstDecl(scope, stmt, isGo)
	case *ast.TypeDecl:
		WalkTypeDecl(scope, stmt, isGo)
	case *ast.VarDecl:
		WalkVarDecl(scope, stmt, isGo)
	case *ast.LabeledStmt:
		WalkStatement(scope, stmt.Statement(), isGo)
	case *ast.SendStmt:
		WalkExpression(scope, stmt.Channel(), isGo)
		WalkExpression(scope, stmt.Expression(), isGo)
	case *ast.IncDecStmt:
		WalkExpression(scope, stmt.Expression(), isGo)
	case *ast.Assignment:
		WalkExpressionList(scope, stmt.ExpressionListLeft(), isGo)
		WalkExpressionList(scope, stmt.ExpressionListRight(), isGo)
	case *ast.Expression:
		WalkExpression(scope, stmt, isGo)
	case *ast.ShortValDecl:
		WalkShortValDecl(scope, stmt, isGo)
	case *ast.GoStmt:
		WalkExpression(scope, stmt.Expression(), true)
	case *ast.ReturnStmt:
		WalkExpressionList(scope, stmt.ExpressionList(), isGo)
	case *ast.BreakStmt: // don't need to do anything
	case *ast.ContinueStmt: // don't need to do anything
	case *ast.GotoStmt: // don't need to do anything
	case *ast.FallThroughStmt: // don't need to do anything
	case *ast.Block:
		WalkBlock(scope, stmt, true, isGo)
	case *ast.IfStmt:
		WalkIfStmt(scope, stmt, isGo)
	case *ast.ExprSwitchStmt:
		newScope := scope.NewChildScope()
		WalkStatement(newScope, stmt.SimpleStmt(), isGo)
		WalkExpression(newScope, stmt.Expression(), isGo)
		for _, caseClause := range stmt.ExprCaseClauses() {
			WalkExprCaseClause(scope, caseClause, isGo)
		}
	case *ast.TypeSwitchStmt:
		newScope := scope.NewChildScope()
		WalkStatement(newScope, stmt.SimpleStmt(), isGo)
		WalkTypeSwitchGuard(newScope, stmt.TypeSwitchGuard(), isGo)
		for _, typeCaseClause := range stmt.TypeCaseClauses() {
			WalkTypeCaseClause(scope, typeCaseClause, isGo)
		}
	case *ast.SelectStmt:
		for _, clause := range stmt.CommClauses() {
			WalkCommClause(scope, clause, isGo)
		}
	case *ast.ForStmt:
		newScope := scope.NewChildScope()
		WalkExpression(newScope, stmt.Expression(), isGo)
		WalkForClause(newScope, stmt.ForClause(), isGo)
		WalkRangeClause(newScope, stmt.RangeClause(), isGo)
		WalkBlock(newScope, stmt.Block(), false, isGo)
	case *ast.DeferStmt:
		WalkExpression(scope, stmt.Expression(), isGo)
	default:
		panic(statement.String())
	}
}

func WalkRangeClause(scope *Scope, rangeClause *ast.RangeClause, isGo bool) {
	if rangeClause == nil {
		return
	}
	WalkExpressionList(scope, rangeClause.ExpressionList(), isGo)
	if rangeClause.IdentifierList() != nil {
		for _, id := range rangeClause.IdentifierList().Identifiers() {
			scope.AddName(id, NewName(id, nil, VARIABLE, ""))
		}
	}
	WalkExpression(scope, rangeClause.Expression(), isGo)
}

func WalkForClause(scope *Scope, forClause *ast.ForClause, isGo bool) {
	if forClause == nil {
		return
	}
	WalkStatement(scope, forClause.InitStmt(), isGo)
	WalkExpression(scope, forClause.Expression(), isGo)
	WalkStatement(scope, forClause.PostStmt(), isGo)
}

func WalkCommClause(scope *Scope, commClause *ast.CommClause, isGo bool) {
	if commClause == nil {
		return
	}
	scope = scope.NewChildScope()
	WalkCommCase(scope, commClause.CommCase(), isGo)
	WalkStatementList(scope, commClause.StatementList(), isGo)
}

func WalkStatementList(scope *Scope, statementList *ast.StatementList, isGo bool) {
	if statementList == nil {
		return
	}
	for _, statement := range statementList.Statements() {
		WalkStatement(scope, statement, isGo)
	}
}

func WalkCommCase(scope *Scope, commCase *ast.CommCase, isGo bool) {
	if commCase == nil {
		return
	}
	WalkSendStmt(scope, commCase.SendStmt(), isGo)
	WalkRecvStmt(scope, commCase.RecvStmt(), isGo)
}

func WalkSendStmt(scope *Scope, stmt *ast.SendStmt, isGo bool) {
	if stmt == nil {
		return
	}
	WalkExpression(scope, stmt.Channel(), isGo)
	WalkExpression(scope, stmt.Expression(), isGo)
}

func WalkRecvStmt(scope *Scope, stmt *ast.RecvStmt, isGo bool) {
	if stmt == nil {
		return
	}
	WalkExpressionList(scope, stmt.ExpressionList(), isGo)
	if stmt.IdentifierList() != nil {
		for _, id := range stmt.IdentifierList().Identifiers() {
			scope.AddName(id, NewName(id, nil, VARIABLE, ""))
		}
	}
	WalkExpression(scope, stmt.RecvExpr(), isGo)
}

func WalkExprCaseClause(scope *Scope, exprCaseClause *ast.ExprCaseClause, isGo bool) {
	if exprCaseClause == nil {
		return
	}
	scope = scope.NewChildScope()
	WalkExprSwitchCase(scope, exprCaseClause.ExprSwitchCase(), isGo)
	WalkStatementList(scope, exprCaseClause.StatementList(), isGo)
}

func WalkExprSwitchCase(scope *Scope, switchCase *ast.ExprSwitchCase, isGo bool) {
	if switchCase == nil {
		return
	}
	WalkExpressionList(scope, switchCase.ExpressionList(), isGo)
}

func WalkTypeCaseClause(scope *Scope, exprCaseClause *ast.TypeCaseClause, isGo bool) {
	if exprCaseClause == nil {
		return
	}
	scope = scope.NewChildScope()
	WalkTypeSwitchCase(scope, exprCaseClause.TypeSwitchCase(), isGo)
	WalkStatementList(scope, exprCaseClause.StatementList(), isGo)
}

func WalkTypeSwitchCase(scope *Scope, switchCase *ast.TypeSwitchCase, isGo bool) {
	if switchCase == nil {
		return
	}
	WalkTypeList(scope, switchCase.TypeList(), isGo)
}

func WalkTypeList(scope *Scope, typeList *ast.TypeList, isGo bool) {
	if typeList == nil {
		return
	}
	for _, type_ := range typeList.Types() {
		WalkType_(scope, type_, isGo)
	}
}

func WalkType_(scope *Scope, type_ *ast.Type_, isGo bool) {
	if type_ == nil {
		return
	}
	// TODO 好像并不需要，走走再看
}

func WalkTypeName(scope *Scope, name *ast.TypeName, isGo bool) {

}

func WalkTypeSwitchGuard(scope *Scope, typeSwitchGuard *ast.TypeSwitchGuard, isGo bool) {
	if typeSwitchGuard == nil {
		return
	}
	WalkPrimaryExpr(scope, typeSwitchGuard.PrimaryExpr(), isGo)
}

func WalkPrimaryExpr(scope *Scope, expr *ast.PrimaryExpr, isGo bool) {
	if expr == nil {
		return
	}
	WalkOperand(scope, expr.Operand(), isGo)
	WalkConversion(scope, expr.Conversion(), isGo)
	WalkMethodExpr(scope, expr.MethodExpr(), isGo)
	WalkPrimaryExpr(scope, expr.PrimaryExpr(), isGo)
	WalkIndex(scope, expr.Index(), isGo)
	WalkSlice(scope, expr.Slice(), isGo)
	WalkTypeAssertion(scope, expr.TypeAssertion(), isGo)
	WalkArguments(scope, expr.Arguments(), isGo)
}

func WalkOperand(scope *Scope, stmt *ast.Operand, isGo bool) {

}

func WalkConversion(scope *Scope, stmt *ast.Conversion, isGo bool) {

}

func WalkMethodExpr(scope *Scope, stmt *ast.MethodExpr, isGo bool) {

}

func WalkIndex(scope *Scope, stmt *ast.Index, isGo bool) {

}

func WalkSlice(scope *Scope, stmt *ast.Slice, isGo bool) {

}

func WalkTypeAssertion(scope *Scope, stmt *ast.TypeAssertion, isGo bool) {

}

func WalkArguments(scope *Scope, stmt *ast.Arguments, isGo bool) {

}

func WalkIfStmt(scope *Scope, stmt *ast.IfStmt, isGo bool) {
	newScope := scope.NewChildScope()
	WalkStatement(newScope, stmt.SimpleStmt(), isGo)
	WalkExpression(newScope, stmt.Expression(), isGo)
	WalkBlock(newScope, stmt.Block(), true, isGo)
	WalkIfStmt(newScope, stmt.Elseif(), isGo)
	WalkBlock(newScope, stmt.Elseblock(), true, isGo)
}

func WalkShortValDecl(scope *Scope, shortValDecl *ast.ShortValDecl, isGo bool) {

}

func WalkExpressionList(scope *Scope, expressionList *ast.ExpressionList, isGo bool) {

}

func WalkExpression(scope *Scope, expression *ast.Expression, isGo bool) {

}
