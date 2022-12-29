package provider

import (
	"github.com/skhoroshavin/automap/internal/mapper/node"
	"github.com/skhoroshavin/automap/internal/mapper/types"
)

func NewMock(name string, children ...Provider) *Mock {
	return &Mock{
		name:     name,
		children: children,
	}
}

func NewMockRequest(name string) *Request {
	return &Request{
		Name: name,
		Type: &types.Type{Name: "Mock"},
	}
}

type Mock struct {
	name     string
	children List
}

func (m *Mock) Signature() *Signature {
	return &Signature{
		Name: Matcher{m.name},
		Type: &types.Type{Name: "Mock"},
	}
}

func (m *Mock) Parent() Provider {
	return nil
}

func (m *Mock) Children() []Provider {
	return m.children
}

func (m *Mock) Dependencies() []Request {
	return nil
}

func (m *Mock) Map(_ node.Node, _ []node.Node) node.Node {
	return node.NewValue(m.name)
}
