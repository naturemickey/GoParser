package ast

type ReceiverType struct {
	BaseNode
}

func (s *ReceiverType) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ReceiverType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ReceiverType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ReceiverType)(nil)
