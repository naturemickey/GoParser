package ast

type GoStmt struct {
	BaseNode
}

func (s *GoStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *GoStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *GoStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*GoStmt)(nil)
