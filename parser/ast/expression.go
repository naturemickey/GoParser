package ast

type Expression struct {
	BaseNode
}

func (s *Expression) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) _SimpleStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) _ExpressionStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) String() string {
	return s.codeBuilder().String()
}

var _ ExpressionStmt = (*Expression)(nil)
