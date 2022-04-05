package ast

type RecvStmt struct {
	BaseNode
}

func (s *RecvStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *RecvStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *RecvStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*RecvStmt)(nil)
