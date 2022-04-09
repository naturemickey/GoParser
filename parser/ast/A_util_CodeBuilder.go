package ast

import (
	"reflect"
	"strings"
)

type CodeBuilder struct {
	ss []string
}

func NewCodeBuilder() *CodeBuilder {
	return &CodeBuilder{}
}

func (cb *CodeBuilder) appendString(s string) *CodeBuilder {
	cb.ss = append(cb.ss, s)
	return cb
}

func (cb *CodeBuilder) appendNode(n INode) *CodeBuilder {
	if n == nil || reflect.ValueOf(n).IsNil() {
		return cb
	}
	for _, s := range n.codeBuilder().ss {
		cb.ss = append(cb.ss, s)
	}
	return cb
}

func (cb *CodeBuilder) newLine() *CodeBuilder {
	return cb.appendString("\n")
}

func (cb *CodeBuilder) tab() *CodeBuilder {
	return cb.appendString("\t")
}

func (cb *CodeBuilder) blank() *CodeBuilder {
	return cb.appendString(" ")
}

func (cb *CodeBuilder) deleteLast() *CodeBuilder {
	cb.ss = cb.ss[:len(cb.ss)-1]
	return cb
}

func (cb *CodeBuilder) String() string {
	return strings.Join(cb.ss, "")
}
