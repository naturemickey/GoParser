package ast

type MethodDecl struct {
	BaseNode
	annotation string
	receiver   Receiver
	funName    string
	signature  *Signature
	block      *Block
}

func NewMethodDecl(annotation string, receiver Receiver, funName string, signature *Signature, block *Block) *MethodDecl {
	return &MethodDecl{receiver: receiver, funName: funName, signature: signature, block: block}
}

func (s *MethodDecl) Annotation() string {
	return s.annotation
}

func (s *MethodDecl) SetAnnotation(annotation string) {
	s.annotation = annotation
}

func (s *MethodDecl) Receiver() Receiver {
	return s.receiver
}

func (s *MethodDecl) SetReceiver(receiver Receiver) {
	s.receiver = receiver
}

func (s *MethodDecl) FunName() string {
	return s.funName
}

func (s *MethodDecl) SetFunName(funName string) {
	s.funName = funName
}

func (s *MethodDecl) Signature() *Signature {
	return s.signature
}

func (s *MethodDecl) SetSignature(signature *Signature) {
	s.signature = signature
}

func (s *MethodDecl) Block() *Block {
	return s.block
}

func (s *MethodDecl) SetBlock(block *Block) {
	s.block = block
}

func (s *MethodDecl) _IFunctionMethodDeclaration_() {
	//TODO implement me
	panic("implement me")
}

func (s *MethodDecl) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("func ").appendNode(s.receiver).blank().appendString(s.funName).appendNode(s.signature)
	if s.block != nil {
		cb.appendNode(s.block)
	}
	return cb
}

func (s *MethodDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *MethodDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*MethodDecl)(nil)
var _ IFunctionMethodDeclaration = (*MethodDecl)(nil)
