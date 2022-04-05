package ast

type ImportPath struct {
	BaseNode
	path string
}

func (s *ImportPath) Path() string {
	return s.path
}

func (s *ImportPath) SetPath(path string) {
	s.path = path
}

func (s *ImportPath) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString(s.path)
}

func (s *ImportPath) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ImportPath) String() string {
	return s.path
}

var _ INode = (*ImportPath)(nil)
