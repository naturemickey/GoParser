package ast

type ShortValDecl struct {
	BaseNode
}

func (s *ShortValDecl) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *ShortValDecl) _SimpleStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *ShortValDecl) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ShortValDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ShortValDecl) String() string {
	return s.codeBuilder().String()
}

var _ SimpleStmt = (*ShortValDecl)(nil)
