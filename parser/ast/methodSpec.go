package ast

type MethodSpec struct {
	BaseNode

	methodName string
	parameters *Parameters
	result     Result
}

func NewMethodSpec(methodName string, parameters *Parameters, result Result) *MethodSpec {
	return &MethodSpec{methodName: methodName, parameters: parameters, result: result}
}

func (s *MethodSpec) MethodName() string {
	return s.methodName
}

func (s *MethodSpec) SetMethodName(methodName string) {
	s.methodName = methodName
}

func (s *MethodSpec) Parameters() *Parameters {
	return s.parameters
}

func (s *MethodSpec) SetParameters(parameters *Parameters) {
	s.parameters = parameters
}

func (s *MethodSpec) Result() Result {
	return s.result
}

func (s *MethodSpec) SetResult(result Result) {
	s.result = result
}

func (s *MethodSpec) _IMethodSpecOrTypeName_() {
	//TODO implement me
	panic("implement me")
}

func (s *MethodSpec) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString(s.methodName).appendNode(s.parameters)
	if s.result != nil {
		cb.blank().appendNode(s.result)
	}
	return cb
}

func (s *MethodSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *MethodSpec) String() string {
	return s.codeBuilder().String()
}

var _ IMethodSpecOrTypeName = (*MethodSpec)(nil)
