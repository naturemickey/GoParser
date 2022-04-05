package ast

type IfStmt struct {
	BaseNode
	expression *Expression
	simpleStmt SimpleStmt
	block      *Block
	elseif     *IfStmt
	elseblock  *Block
}

func (s *IfStmt) Expression() *Expression {
	return s.expression
}

func (s *IfStmt) SetExpression(expression *Expression) {
	s.expression = expression
}

func (s *IfStmt) SimpleStmt() SimpleStmt {
	return s.simpleStmt
}

func (s *IfStmt) SetSimpleStmt(simpleStmt SimpleStmt) {
	s.simpleStmt = simpleStmt
}

func (s *IfStmt) Block() *Block {
	return s.block
}

func (s *IfStmt) SetBlock(block *Block) {
	s.block = block
}

func (s *IfStmt) Elseif() *IfStmt {
	return s.elseif
}

func (s *IfStmt) SetElseif(elseif *IfStmt) {
	s.elseif = elseif
}

func (s *IfStmt) Elseblock() *Block {
	return s.elseblock
}

func (s *IfStmt) SetElseblock(elseblock *Block) {
	s.elseblock = elseblock
}

func (s *IfStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *IfStmt) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("if ")
	if s.simpleStmt != nil {
		cb.appendNode(s.simpleStmt).appendString("; ")
	}
	cb.appendNode(s.expression).blank()
	cb.appendNode(s.block)

	if s.elseif != nil {
		cb.appendString(" else ").appendNode(s.elseif)
	} else if s.elseblock != nil {
		cb.appendString(" else ").appendNode(s.elseblock)
	}
	return cb
}

func (s *IfStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *IfStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*IfStmt)(nil)
