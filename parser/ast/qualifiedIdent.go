package ast

type QualifiedIdent struct {
	BaseNode

	id1 string
	id2 string
}

func (s *QualifiedIdent) Id1() string {
	return s.id1
}

func (s *QualifiedIdent) SetId1(id1 string) {
	s.id1 = id1
}

func (s *QualifiedIdent) Id2() string {
	return s.id2
}

func (s *QualifiedIdent) SetId2(id2 string) {
	s.id2 = id2
}

func (s *QualifiedIdent) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString(s.id1).appendString(".").appendString(s.id2)
}

func (s *QualifiedIdent) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *QualifiedIdent) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*QualifiedIdent)(nil)
