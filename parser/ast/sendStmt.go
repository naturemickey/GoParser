package ast

type SendStmt struct {
	BaseNode
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

var _ INode = (*SendStmt)(nil)
