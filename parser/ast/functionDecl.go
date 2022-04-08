package ast

type FunctionDecl struct {
	BaseNode
	funName   string
	signature *Signature
	block     *Block
}

func NewFunctionDecl(funName string, signature *Signature, block *Block) *FunctionDecl {
	return &FunctionDecl{funName: funName, signature: signature, block: block}
}

func (s *FunctionDecl) FunName() string {
	return s.funName
}

func (s *FunctionDecl) SetFunName(funName string) {
	s.funName = funName
}

func (s *FunctionDecl) Signature() *Signature {
	return s.signature
}

func (s *FunctionDecl) SetSignature(signature *Signature) {
	s.signature = signature
}

func (s *FunctionDecl) Block() *Block {
	return s.block
}

func (s *FunctionDecl) SetBlock(block *Block) {
	s.block = block
}

func (s *FunctionDecl) _IFunctionMethodDeclaration_() {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionDecl) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("func ").appendString(s.funName).appendNode(s.signature)
	if s.block != nil {
		cb.appendNode(s.block)
	}
	return cb
}

func (s *FunctionDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*FunctionDecl)(nil)
var _ IFunctionMethodDeclaration = (*FunctionDecl)(nil)
