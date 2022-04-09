package ast

type Signature struct {
	BaseNode

	parameters *Parameters
	result     Result
}

func NewSignature(parameters *Parameters, result Result) *Signature {
	return &Signature{parameters: parameters, result: result}
}

func (s *Signature) Parameters() *Parameters {
	return s.parameters
}

func (s *Signature) SetParameters(parameters *Parameters) {
	s.parameters = parameters
}

func (s *Signature) Result() Result {
	return s.result
}

func (s *Signature) SetResult(result Result) {
	s.result = result
}

func (s *Signature) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendNode(s.parameters)
	if s.result != nil {
		cb.blank().appendNode(s.result)
	}
	return cb
}

func (s *Signature) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Signature) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Signature)(nil)
