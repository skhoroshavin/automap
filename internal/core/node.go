package core

import (
	"strings"
)

type Node interface {
	Build(fn *FuncBody, indent int) string
}

func BuildFuncBody(node Node) *FuncBody {
	res := new(FuncBody)
	res.Result = node.Build(res, 1)
	return res
}

type NamedNode struct {
	Name  string
	Value Node
}

type ExprNode struct {
	Value string
}

func (n *ExprNode) Build(_ *FuncBody, _ int) string {
	return n.Value
}

type StructNode struct {
	Name      string
	Fields    []NamedNode
	IsPointer bool
}

func (n *StructNode) Build(fn *FuncBody, indent int) string {
	s := strings.Builder{}
	if n.IsPointer {
		s.WriteString("&")
	}
	s.WriteString(n.Name)
	s.WriteString("{\n")
	for _, field := range n.Fields {
		writeIndent(&s, indent+1)
		s.WriteString(field.Name)
		s.WriteString(": ")
		s.WriteString(field.Value.Build(fn, indent+1))
		s.WriteString(",\n")
	}
	writeIndent(&s, indent)
	s.WriteString("}")
	return s.String()
}

type FuncNode struct {
	Name string

	FuncName string
	Inputs   []Node
}

func writeIndent(s *strings.Builder, indent int) {
	for i := 0; i != indent; i++ {
		s.WriteString("\t")
	}
}
