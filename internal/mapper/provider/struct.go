package provider

import (
	"github.com/skhoroshavin/automap/internal/mapper/node"
	"github.com/skhoroshavin/automap/internal/mapper/types"
)

func NewStruct(base string, typ *types.Type) *Struct {
	deps := make([]Request, len(typ.Fields))
	for i, field := range typ.Fields {
		deps[i] = Request{
			Name: base + field.Name,
			Type: field.Type,
		}
	}

	return &Struct{
		sig: Signature{
			Name: base,
			Type: typ,
		},
		deps: deps,
	}
}

type Struct struct {
	sig  Signature
	name string
	deps []Request
}

func (s *Struct) Signature() *Signature {
	return &s.sig
}

func (s *Struct) Parent() Provider {
	return nil
}

func (s *Struct) Children() []Provider {
	// TODO: Provide methods?
	return nil
}

func (s *Struct) Dependencies() []Request {
	return s.deps
}

func (s *Struct) Map(_ node.Node, args []node.Node) node.Node {
	// TODO: Add check that we have same number of args as fields
	fields := make([]*node.Field, len(s.sig.Type.Fields))
	for i, field := range s.sig.Type.Fields {
		fields[i] = node.NewField(field.Name, args[i])
	}
	return node.NewStruct(s.sig.Type.Name, fields...)
}
