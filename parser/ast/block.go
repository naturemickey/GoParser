package ast

type Block struct {
	BaseNode
	statementList *StatementList
}

func (s *Block) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *Block) StatementList() *StatementList {
	return s.statementList
}

func (s *Block) SetStatementList(statementList *StatementList) {
	s.statementList = statementList
}

func (s *Block) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("{\n")
	cb.appendNode(s.statementList)
	cb.appendString("}")
	return cb
}

func (s *Block) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Block) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*Block)(nil)
