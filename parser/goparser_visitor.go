package parser

import (
	"GoParser/parser/antlr4"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type GoParserVisitorImpl struct {
	*antlr.BaseParseTreeVisitor
}

func (g *GoParserVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	//TODO implement me
	panic("implement me")
}
func (g *GoParserVisitorImpl) VisitChildren(node antlr.RuleNode) interface{} {
	//TODO implement me
	panic("implement me")
}
func (g *GoParserVisitorImpl) VisitTerminal(node antlr.TerminalNode) interface{} {
	//TODO implement me
	panic("implement me")
}
func (g *GoParserVisitorImpl) VisitErrorNode(node antlr.ErrorNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitSourceFile(ctx *antlr4.SourceFileContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitPackageClause(ctx *antlr4.PackageClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitImportDecl(ctx *antlr4.ImportDeclContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitImportSpec(ctx *antlr4.ImportSpecContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitImportPath(ctx *antlr4.ImportPathContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitDeclaration(ctx *antlr4.DeclarationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitConstDecl(ctx *antlr4.ConstDeclContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitConstSpec(ctx *antlr4.ConstSpecContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitIdentifierList(ctx *antlr4.IdentifierListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitExpressionList(ctx *antlr4.ExpressionListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitTypeDecl(ctx *antlr4.TypeDeclContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitTypeSpec(ctx *antlr4.TypeSpecContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitFunctionDecl(ctx *antlr4.FunctionDeclContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitMethodDecl(ctx *antlr4.MethodDeclContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitReceiver(ctx *antlr4.ReceiverContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitVarDecl(ctx *antlr4.VarDeclContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitVarSpec(ctx *antlr4.VarSpecContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitBlock(ctx *antlr4.BlockContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitStatementList(ctx *antlr4.StatementListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitStatement(ctx *antlr4.StatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitSimpleStmt(ctx *antlr4.SimpleStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitExpressionStmt(ctx *antlr4.ExpressionStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitSendStmt(ctx *antlr4.SendStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitIncDecStmt(ctx *antlr4.IncDecStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitAssignment(ctx *antlr4.AssignmentContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitAssign_op(ctx *antlr4.Assign_opContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitShortVarDecl(ctx *antlr4.ShortVarDeclContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitEmptyStmt(ctx *antlr4.EmptyStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitLabeledStmt(ctx *antlr4.LabeledStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitReturnStmt(ctx *antlr4.ReturnStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitBreakStmt(ctx *antlr4.BreakStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitContinueStmt(ctx *antlr4.ContinueStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitGotoStmt(ctx *antlr4.GotoStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitFallthroughStmt(ctx *antlr4.FallthroughStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitDeferStmt(ctx *antlr4.DeferStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitIfStmt(ctx *antlr4.IfStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitSwitchStmt(ctx *antlr4.SwitchStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitExprSwitchStmt(ctx *antlr4.ExprSwitchStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitExprCaseClause(ctx *antlr4.ExprCaseClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitExprSwitchCase(ctx *antlr4.ExprSwitchCaseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitTypeSwitchStmt(ctx *antlr4.TypeSwitchStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitTypeSwitchGuard(ctx *antlr4.TypeSwitchGuardContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitTypeCaseClause(ctx *antlr4.TypeCaseClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitTypeSwitchCase(ctx *antlr4.TypeSwitchCaseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitTypeList(ctx *antlr4.TypeListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitSelectStmt(ctx *antlr4.SelectStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitCommClause(ctx *antlr4.CommClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitCommCase(ctx *antlr4.CommCaseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitRecvStmt(ctx *antlr4.RecvStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitForStmt(ctx *antlr4.ForStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitForClause(ctx *antlr4.ForClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitRangeClause(ctx *antlr4.RangeClauseContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitGoStmt(ctx *antlr4.GoStmtContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitType_(ctx *antlr4.Type_Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitTypeName(ctx *antlr4.TypeNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitTypeLit(ctx *antlr4.TypeLitContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitArrayType(ctx *antlr4.ArrayTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitArrayLength(ctx *antlr4.ArrayLengthContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitElementType(ctx *antlr4.ElementTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitPointerType(ctx *antlr4.PointerTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitInterfaceType(ctx *antlr4.InterfaceTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitSliceType(ctx *antlr4.SliceTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitMapType(ctx *antlr4.MapTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitChannelType(ctx *antlr4.ChannelTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitMethodSpec(ctx *antlr4.MethodSpecContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitFunctionType(ctx *antlr4.FunctionTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitSignature(ctx *antlr4.SignatureContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitResult(ctx *antlr4.ResultContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitParameters(ctx *antlr4.ParametersContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitParameterDecl(ctx *antlr4.ParameterDeclContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitExpression(ctx *antlr4.ExpressionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitPrimaryExpr(ctx *antlr4.PrimaryExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitConversion(ctx *antlr4.ConversionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitNonNamedType(ctx *antlr4.NonNamedTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitOperand(ctx *antlr4.OperandContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitLiteral(ctx *antlr4.LiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitBasicLit(ctx *antlr4.BasicLitContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitInteger(ctx *antlr4.IntegerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitOperandName(ctx *antlr4.OperandNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitQualifiedIdent(ctx *antlr4.QualifiedIdentContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitCompositeLit(ctx *antlr4.CompositeLitContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitLiteralType(ctx *antlr4.LiteralTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitLiteralValue(ctx *antlr4.LiteralValueContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitElementList(ctx *antlr4.ElementListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitKeyedElement(ctx *antlr4.KeyedElementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitKey(ctx *antlr4.KeyContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitElement(ctx *antlr4.ElementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitStructType(ctx *antlr4.StructTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitFieldDecl(ctx *antlr4.FieldDeclContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitString_(ctx *antlr4.String_Context) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitEmbeddedField(ctx *antlr4.EmbeddedFieldContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitFunctionLit(ctx *antlr4.FunctionLitContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitIndex(ctx *antlr4.IndexContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitSlice(ctx *antlr4.SliceContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitTypeAssertion(ctx *antlr4.TypeAssertionContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitArguments(ctx *antlr4.ArgumentsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitMethodExpr(ctx *antlr4.MethodExprContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitReceiverType(ctx *antlr4.ReceiverTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GoParserVisitorImpl) VisitEos(ctx *antlr4.EosContext) interface{} {
	//TODO implement me
	panic("implement me")
}

var _ antlr4.GoParserVisitor = (*GoParserVisitorImpl)(nil)
