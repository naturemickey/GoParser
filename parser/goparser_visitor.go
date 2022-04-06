package parser

import (
	"GoParser/parser/antlr4"
	"github.com/antlr/antlr4/runtime/Go/antlr"
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
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitPackageClause(ctx *antlr4.PackageClauseContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitImportDecl(ctx *antlr4.ImportDeclContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitImportSpec(ctx *antlr4.ImportSpecContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitImportPath(ctx *antlr4.ImportPathContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitDeclaration(ctx *antlr4.DeclarationContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitConstDecl(ctx *antlr4.ConstDeclContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitConstSpec(ctx *antlr4.ConstSpecContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitIdentifierList(ctx *antlr4.IdentifierListContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitExpressionList(ctx *antlr4.ExpressionListContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitTypeDecl(ctx *antlr4.TypeDeclContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitTypeSpec(ctx *antlr4.TypeSpecContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitFunctionDecl(ctx *antlr4.FunctionDeclContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitMethodDecl(ctx *antlr4.MethodDeclContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitReceiver(ctx *antlr4.ReceiverContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitVarDecl(ctx *antlr4.VarDeclContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitVarSpec(ctx *antlr4.VarSpecContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitBlock(ctx *antlr4.BlockContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitStatementList(ctx *antlr4.StatementListContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitStatement(ctx *antlr4.StatementContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitSimpleStmt(ctx *antlr4.SimpleStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitExpressionStmt(ctx *antlr4.ExpressionStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitSendStmt(ctx *antlr4.SendStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitIncDecStmt(ctx *antlr4.IncDecStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitAssignment(ctx *antlr4.AssignmentContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitAssign_op(ctx *antlr4.Assign_opContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitShortVarDecl(ctx *antlr4.ShortVarDeclContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitEmptyStmt(ctx *antlr4.EmptyStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitLabeledStmt(ctx *antlr4.LabeledStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitReturnStmt(ctx *antlr4.ReturnStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitBreakStmt(ctx *antlr4.BreakStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitContinueStmt(ctx *antlr4.ContinueStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitGotoStmt(ctx *antlr4.GotoStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitFallthroughStmt(ctx *antlr4.FallthroughStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitDeferStmt(ctx *antlr4.DeferStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitIfStmt(ctx *antlr4.IfStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitSwitchStmt(ctx *antlr4.SwitchStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitExprSwitchStmt(ctx *antlr4.ExprSwitchStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitExprCaseClause(ctx *antlr4.ExprCaseClauseContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitExprSwitchCase(ctx *antlr4.ExprSwitchCaseContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitTypeSwitchStmt(ctx *antlr4.TypeSwitchStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitTypeSwitchGuard(ctx *antlr4.TypeSwitchGuardContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitTypeCaseClause(ctx *antlr4.TypeCaseClauseContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitTypeSwitchCase(ctx *antlr4.TypeSwitchCaseContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitTypeList(ctx *antlr4.TypeListContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitSelectStmt(ctx *antlr4.SelectStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitCommClause(ctx *antlr4.CommClauseContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitCommCase(ctx *antlr4.CommCaseContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitRecvStmt(ctx *antlr4.RecvStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitForStmt(ctx *antlr4.ForStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitForClause(ctx *antlr4.ForClauseContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitRangeClause(ctx *antlr4.RangeClauseContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitGoStmt(ctx *antlr4.GoStmtContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitType_(ctx *antlr4.Type_Context) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitTypeName(ctx *antlr4.TypeNameContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitTypeLit(ctx *antlr4.TypeLitContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitArrayType(ctx *antlr4.ArrayTypeContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitArrayLength(ctx *antlr4.ArrayLengthContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitElementType(ctx *antlr4.ElementTypeContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitPointerType(ctx *antlr4.PointerTypeContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitInterfaceType(ctx *antlr4.InterfaceTypeContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitSliceType(ctx *antlr4.SliceTypeContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitMapType(ctx *antlr4.MapTypeContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitChannelType(ctx *antlr4.ChannelTypeContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitMethodSpec(ctx *antlr4.MethodSpecContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitFunctionType(ctx *antlr4.FunctionTypeContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitSignature(ctx *antlr4.SignatureContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitResult(ctx *antlr4.ResultContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitParameters(ctx *antlr4.ParametersContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitParameterDecl(ctx *antlr4.ParameterDeclContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitExpression(ctx *antlr4.ExpressionContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitPrimaryExpr(ctx *antlr4.PrimaryExprContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitConversion(ctx *antlr4.ConversionContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitNonNamedType(ctx *antlr4.NonNamedTypeContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitOperand(ctx *antlr4.OperandContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitLiteral(ctx *antlr4.LiteralContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitBasicLit(ctx *antlr4.BasicLitContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitInteger(ctx *antlr4.IntegerContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitOperandName(ctx *antlr4.OperandNameContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitQualifiedIdent(ctx *antlr4.QualifiedIdentContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitCompositeLit(ctx *antlr4.CompositeLitContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitLiteralType(ctx *antlr4.LiteralTypeContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitLiteralValue(ctx *antlr4.LiteralValueContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitElementList(ctx *antlr4.ElementListContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitKeyedElement(ctx *antlr4.KeyedElementContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitKey(ctx *antlr4.KeyContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitElement(ctx *antlr4.ElementContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitStructType(ctx *antlr4.StructTypeContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitFieldDecl(ctx *antlr4.FieldDeclContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitString_(ctx *antlr4.String_Context) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitEmbeddedField(ctx *antlr4.EmbeddedFieldContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitFunctionLit(ctx *antlr4.FunctionLitContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitIndex(ctx *antlr4.IndexContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitSlice(ctx *antlr4.SliceContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitTypeAssertion(ctx *antlr4.TypeAssertionContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitArguments(ctx *antlr4.ArgumentsContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitMethodExpr(ctx *antlr4.MethodExprContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitReceiverType(ctx *antlr4.ReceiverTypeContext) interface{} {
	panic("implement me")
}

func (visitor *GoParserVisitorImpl) VisitEos(ctx *antlr4.EosContext) interface{} {
	panic("implement me")
}

var _ antlr4.GoParserVisitor = (*GoParserVisitorImpl)(nil)
