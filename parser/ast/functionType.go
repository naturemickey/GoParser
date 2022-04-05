package ast

type FunctionType struct {
	BaseNode

	signature *Signature
}

func (s *FunctionType) Signature() *Signature {
	return s.signature
}

func (s *FunctionType) SetSignature(signature *Signature) {
	s.signature = signature
}

func (s *FunctionType) _TypeLit_() {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionType) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("func ").appendNode(s.signature)
}

func (s *FunctionType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionType) String() string {
	return s.codeBuilder().String()
}

var _ TypeLit = (*FunctionType)(nil)
