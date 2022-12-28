package mapper

import "github.com/skhoroshavin/automap/internal/mapper/node"

// Provider represents single provider
type Provider interface {
	Name() string

	Match(Request) bool

	Parent() Provider
	Children() []Provider

	Dependencies() []Request

	Map(node.Node, []node.Node) node.Node
}

type MockProvider struct {
	name     string
	children []Provider
}

func (m *MockProvider) Name() string {
	return m.name
}

func (m *MockProvider) Match(request Request) bool {
	return request.Name == m.name
}

func (m *MockProvider) Parent() Provider {
	return nil
}

func (m *MockProvider) Children() []Provider {
	return m.children
}

func (m *MockProvider) Dependencies() []Request {
	return nil
}

func (m *MockProvider) Map(base node.Node, args []node.Node) node.Node {
	return node.NewValue(m.name)
}

type VarProvider struct {
	name  string
	value string
	typ   OldType

	children []Provider
}

type FieldProvider struct {
	name  string
	value string
	typ   OldType

	children []Provider
}

type MethodProvider struct {
	name  string
	value string
	deps  []Request
}

// Nodes
// Node ID
// Node name
// Node Type
//  * Ident (name)
//  * Field (base.name)
//  * Struct (name{fields...})
//  * Method (base.name(args...))
//  * Func (name(args...))
//  * Addr (&base)
//  * Deref (*base)
// Node deps
//  * Base
//  * Args...

// Providers
// Provider name
// Provider Node name
// Provider Node type
