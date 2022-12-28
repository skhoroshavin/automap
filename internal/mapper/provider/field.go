package provider

import (
	"github.com/skhoroshavin/automap/internal/mapper/ast"
	"github.com/skhoroshavin/automap/internal/mapper/node"
	"github.com/skhoroshavin/automap/internal/mapper/types"
)

func NewField(parent Provider, field *types.Var) *Field {
	return &Field{
		sig: Signature{
			Name: parent.Signature().Name + field.Name,
			Type: field.Type,
		},
		parent: parent,
		name:   field.Name,
	}
}

type Field struct {
	sig    Signature
	parent Provider
	name   string
}

func (p *Field) Signature() *Signature {
	return &p.sig
}

func (p *Field) Parent() Provider {
	return p.parent
}

func (p *Field) Children() []Provider {
	return p.sig.Unpack(p)
}

func (p *Field) Dependencies() []Request {
	return nil
}

func (p *Field) Map(base node.Node, _ []node.Node) node.Node {
	// TODO: Remove this hack and implement proper mapper
	mapper := new(ast.Mapper)
	base.CompileTo(mapper)
	value := mapper.Result.(*ast.ValueExpr).Value
	return node.NewValue(value + "." + p.name)
}
