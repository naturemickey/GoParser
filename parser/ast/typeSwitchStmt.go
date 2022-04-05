package ast

type TypeSwitchStmt struct {
	BaseNode
	typeSwitchGuard *TypeSwitchGuard
	simpleStmt      SimpleStmt
	typeCaseClauses []*TypeCaseClause
}

func (s *TypeSwitchStmt) TypeSwitchGuard() *TypeSwitchGuard {
	return s.typeSwitchGuard
}

func (s *TypeSwitchStmt) SetTypeSwitchGuard(typeSwitchGuard *TypeSwitchGuard) {
	s.typeSwitchGuard = typeSwitchGuard
}

func (s *TypeSwitchStmt) SimpleStmt() SimpleStmt {
	return s.simpleStmt
}

func (s *TypeSwitchStmt) SetSimpleStmt(simpleStmt SimpleStmt) {
	s.simpleStmt = simpleStmt
}

func (s *TypeSwitchStmt) TypeCaseClauses() []*TypeCaseClause {
	return s.typeCaseClauses
}

func (s *TypeSwitchStmt) SetTypeCaseClauses(typeCaseClauses []*TypeCaseClause) {
	s.typeCaseClauses = typeCaseClauses
}

func (s *TypeSwitchStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSwitchStmt) _SwitchStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSwitchStmt) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("switch ")
	if s.simpleStmt != nil {
		cb.appendNode(s.simpleStmt).appendString("; ")
	}
	cb.appendNode(s.typeSwitchGuard)
	cb.appendString("{\n")
	for _, clause := range s.typeCaseClauses {
		cb.appendNode(clause).newLine()
	}
	cb.appendString("}")
	return cb
}

func (s *TypeSwitchStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSwitchStmt) String() string {
	return s.codeBuilder().String()
}

var _ SwitchStmt = (*TypeSwitchStmt)(nil)
