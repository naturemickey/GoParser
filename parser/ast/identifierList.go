package ast

type IdentifierList struct {
	BaseNode
	identifiers []string
}

func (s *IdentifierList) Identifiers() []string {
	return s.identifiers
}

func (s *IdentifierList) SetIdentifiers(identifiers []string) {
	s.identifiers = identifiers
}

func (s *IdentifierList) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	for _, identifier := range s.identifiers {
		cb.appendString(identifier).appendString(", ")
	}
	cb.deleteLast()
	return cb
}

func (s *IdentifierList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *IdentifierList) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*IdentifierList)(nil)
