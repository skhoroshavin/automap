package mapper

import "github.com/skhoroshavin/automap/internal/mapper/node"

type Request struct {
	Name string
	Type Type
}

type Provider interface {
	Match(Request) bool
	Dependencies() []Request
	Resolve() (node.Node, error)
}
