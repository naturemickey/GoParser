package ast

type SendStmt struct {
	BaseNode
	channel    *Expression
	expression *Expression
}

func NewSendStmt(channel *Expression, expression *Expression) *SendStmt {
	return &SendStmt{channel: channel, expression: expression}
}

func (s *SendStmt) Channel() *Expression {
	return s.channel
}

func (s *SendStmt) SetChannel(channel *Expression) {
	s.channel = channel
}

func (s *SendStmt) Expression() *Expression {
	return s.expression
}

func (s *SendStmt) SetExpression(expression *Expression) {
	s.expression = expression
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
	return NewCodeBuilder().appendNode(s.channel).appendString(" <- ").appendNode(s.expression)
}

func (s *SendStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *SendStmt) String() string {
	return s.codeBuilder().String()
}

var _ SimpleStmt = (*SendStmt)(nil)
