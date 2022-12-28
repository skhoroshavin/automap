package provider

import (
	"github.com/skhoroshavin/automap/internal/mapper/ast"
	"github.com/skhoroshavin/automap/internal/mapper/node"
	"github.com/skhoroshavin/automap/internal/mapper/types"
)

func NewMethod(parent Provider, method *types.Func) *Method {
	deps := make([]Request, len(method.Args))
	for i, arg := range method.Args {
		deps[i] = Request{
			Name: parent.Signature().Name + arg.Name,
			Type: arg.Type,
		}
	}

	return &Method{
		sig: Signature{
			Name: parent.Signature().Name + method.Name,
			Type: method.ReturnType,
		},
		parent: parent,
		name:   method.Name,
		deps:   deps,
	}
}

type Method struct {
	sig    Signature
	parent Provider
	name   string
	deps   []Request
}

func (p *Method) Signature() *Signature {
	return &p.sig
}

func (p *Method) Parent() Provider {
	return p.parent
}

func (p *Method) Children() []Provider {
	return p.sig.Unpack(p)
}

func (p *Method) Dependencies() []Request {
	return p.deps
}

func (p *Method) Map(base node.Node, _ []node.Node) node.Node {
	// TODO: Remove this hack and implement proper mapper
	mapper := new(ast.Mapper)
	base.CompileTo(mapper)
	value := mapper.Result.(*ast.ValueExpr).Value
	return node.NewValue(value + "." + p.name + "()")
}
