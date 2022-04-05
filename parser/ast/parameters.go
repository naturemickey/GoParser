package ast

type Parameters struct {
	BaseNode

	parameterDecls []*ParameterDecl
}

func (s *Parameters) ParameterDecls() []*ParameterDecl {
	return s.parameterDecls
}

func (s *Parameters) SetParameterDecls(parameterDecls []*ParameterDecl) {
	s.parameterDecls = parameterDecls
}

func (s *Parameters) _Result_() {
	//TODO implement me
	panic("implement me")
}

func (s *Parameters) _Receiver_() {
	//TODO implement me
	panic("implement me")
}

func (s *Parameters) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("(")
	for _, decl := range s.parameterDecls {
		cb.appendNode(decl).appendString(", ")
	}
	cb.deleteLast()

	cb.appendString(")")
	return cb
}

func (s *Parameters) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Parameters) String() string {
	return s.codeBuilder().String()
}

var _ Receiver = (*Parameters)(nil)
var _ Result = (*Parameters)(nil)
