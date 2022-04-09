package semantic

type Scope struct {
	names  map[string]*Name
	parent *Scope
}

func NewScopeRoot() *Scope {
	return new(Scope)
}

func (s *Scope) AddName(name string, val *Name) {
	s.names[name] = val
}

func (s *Scope) NewChildScope() *Scope {
	child := new(Scope)
	child.parent = s
	return child
}

func (s *Scope) popSelf() *Scope {
	if s.parent != nil {
		return s.parent
	}
	panic("you have a bug.")
}

func (s *Scope) isRoot() bool {
	return s.parent == nil
}
