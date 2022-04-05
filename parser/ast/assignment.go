package ast

type Assignment struct {
	BaseNode
}

func (s *Assignment) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *Assignment) _SimpleStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *Assignment) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Assignment) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Assignment) String() string {
	return s.codeBuilder().String()
}

var _ SimpleStmt = (*Assignment)(nil)
