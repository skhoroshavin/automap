package provider

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/node"
	"github.com/skhoroshavin/automap/internal/mapper/types"
)

type Request struct {
	Name string
	Type *types.Type
}

func (r *Request) Resolve(providers List) (res node.Node, err error) {
	append(providers, r.TypeCasts()...).ForEach(func(p Provider) bool {
		if !p.Signature().Match(r) {
			return false
		}

		res, err = compile(p, providers)
		if err != nil {
			return false
		}

		return true
	})

	if res == nil && err == nil {
		err = fmt.Errorf("failed to resolve request %s %s", r.Name, r.Type.ID())
	}
	return
}

func (r *Request) TypeCasts() (res List) {
	if r.Type.IsPointer {
		// Add deref provider
	} else {
		// Add ref provider
	}

	if r.Type.IsStruct {
		// Add struct builder provider
	}

	return
}

func compile(p Provider, providers List) (res node.Node, err error) {
	if p == nil {
		return nil, nil
	}

	base, err := compile(p.Parent(), providers)
	if err != nil {
		return
	}

	deps := make([]node.Node, len(p.Dependencies()))
	for i, dep := range p.Dependencies() {
		deps[i], err = dep.Resolve(providers)
		if err != nil {
			return
		}
	}

	res = p.Map(base, deps)
	return
}
