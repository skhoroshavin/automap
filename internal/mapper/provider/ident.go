package provider

import (
	"github.com/skhoroshavin/automap/internal/mapper/node"
	"github.com/skhoroshavin/automap/internal/mapper/types"
)

func NewIdent(name string, typ *types.Type) *Ident {
	return &Ident{
		sig: Signature{
			Name: Matcher{name},
			Type: typ,
		},
		name: name,
	}
}

type Ident struct {
	sig  Signature
	name string
}

func (p *Ident) Signature() *Signature {
	return &p.sig
}

func (p *Ident) Parent() Provider {
	return nil
}

func (p *Ident) Children() []Provider {
	return p.sig.Unpack(p)
}

func (p *Ident) Dependencies() []Request {
	return nil
}

func (p *Ident) Map(_ node.Node, _ []node.Node) node.Node {
	return node.NewValue(p.name)
}
