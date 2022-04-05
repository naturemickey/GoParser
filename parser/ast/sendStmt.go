package ast

type SendStmt struct {
	BaseNode
}

func (s *SendStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *SendStmt) _SimpleStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *SendStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *SendStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *SendStmt) String() string {
	return s.codeBuilder().String()
}

var _ SimpleStmt = (*SendStmt)(nil)
