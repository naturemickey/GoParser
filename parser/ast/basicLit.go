package ast

type BasicLit struct {
	BaseNode

	isNil   bool
	integer *Integer
	string_ string
	float   string
}

func NewBasicLit(isNil bool, integer *Integer, string_ string, float string) *BasicLit {
	return &BasicLit{isNil: isNil, integer: integer, string_: string_, float: float}
}

func (s *BasicLit) IsNil() bool {
	return s.isNil
}

func (s *BasicLit) SetIsNil(isNil bool) {
	s.isNil = isNil
}

func (s *BasicLit) Integer() *Integer {
	return s.integer
}

func (s *BasicLit) SetInteger(integer *Integer) {
	s.integer = integer
}

func (s *BasicLit) String_() string {
	return s.string_
}

func (s *BasicLit) SetString_(string_ string) {
	s.string_ = string_
}

func (s *BasicLit) Float() string {
	return s.float
}

func (s *BasicLit) SetFloat(float string) {
	s.float = float
}

func (s *BasicLit) _Literal_() {
	//TODO implement me
	panic("implement me")
}

func (s *BasicLit) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	if s.isNil {
		return cb.appendString("nil")
	}
	if s.integer != nil {
		return cb.appendNode(s.integer)
	}
	if s.string_ != "" {
		return cb.appendString(s.string_)
	}
	return cb.appendString(s.float)
}

func (s *BasicLit) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *BasicLit) String() string {
	return s.codeBuilder().String()
}

var _ Literal = (*BasicLit)(nil)
