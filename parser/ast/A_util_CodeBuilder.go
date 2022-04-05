package ast

import "strings"

type CodeBuilder struct {
	ss []string
}

func (cb *CodeBuilder) appendString(s string) *CodeBuilder {
	cb.ss = append(cb.ss, s)
	return cb
}

func (cb *CodeBuilder) appendNode(n INode) *CodeBuilder {
	for _, s := range n.codeBuilder().ss {
		cb.ss = append(cb.ss, s)
	}
	return cb
}

func (cb *CodeBuilder) newLine() *CodeBuilder {
	return cb.appendString("\n")
}

func (cb *CodeBuilder) String() string {
	return strings.Join(cb.ss, "")
}
