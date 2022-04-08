package ast

type CommCase struct {
	BaseNode

	sendStmt *SendStmt
	recvStmt *RecvStmt
}

func NewCommCase(sendStmt *SendStmt, recvStmt *RecvStmt) *CommCase {
	return &CommCase{sendStmt: sendStmt, recvStmt: recvStmt}
}

func (s *CommCase) SendStmt() *SendStmt {
	return s.sendStmt
}

func (s *CommCase) SetSendStmt(sendStmt *SendStmt) {
	s.sendStmt = sendStmt
}

func (s *CommCase) RecvStmt() *RecvStmt {
	return s.recvStmt
}

func (s *CommCase) SetRecvStmt(recvStmt *RecvStmt) {
	s.recvStmt = recvStmt
}

func (s *CommCase) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	if s.sendStmt != nil {
		cb.appendString("case ").appendNode(s.sendStmt)
	} else if s.recvStmt != nil {
		cb.appendString("case ").appendNode(s.recvStmt)
	} else {
		cb.appendString("default")
	}
	return cb
}

func (s *CommCase) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *CommCase) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*CommCase)(nil)
