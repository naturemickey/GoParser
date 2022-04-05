package ast

type SelectStmt struct {
	BaseNode
	commClauses []*CommClause
}

func (s *SelectStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *SelectStmt) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("select {\n")
	for _, clause := range s.commClauses {
		cb.appendNode(clause).newLine()
	}
	cb.appendString("}")
	return cb
}

func (s *SelectStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *SelectStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*SelectStmt)(nil)
