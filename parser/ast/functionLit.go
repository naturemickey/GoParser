package ast

type FunctionLit struct {
	BaseNode

	signature *Signature
	block     *Block
}

func (s *FunctionLit) Signature() *Signature {
	return s.signature
}

func (s *FunctionLit) SetSignature(signature *Signature) {
	s.signature = signature
}

func (s *FunctionLit) Block() *Block {
	return s.block
}

func (s *FunctionLit) SetBlock(block *Block) {
	s.block = block
}

func (s *FunctionLit) _Literal_() {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionLit) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("func").appendNode(s.signature).appendNode(s.block)
}

func (s *FunctionLit) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionLit) String() string {
	return s.codeBuilder().String()
}

var _ Literal = (*FunctionLit)(nil)
