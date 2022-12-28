package mapper

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/node"
	"github.com/skhoroshavin/automap/internal/mapper/provider"
)

// Context represents mapping context
type Context struct {
	providers provider.List
}

// AddProvider adds provider to the list of available providers
func (c *Context) AddProvider(p provider.Provider) {
	c.providers = append(c.providers, p)
}

// Resolve tries to resolve given Request into node tree
func (c *Context) Resolve(req *provider.Request) (res node.Node, err error) {
	providers := append(c.providers, req.TypeCasts()...)
	providers.ForEach(func(p provider.Provider) bool {
		if !p.Signature().Match(req) {
			return false
		}

		res, err = c.compile(p)
		if err != nil {
			return false
		}

		return true
	})

	if res == nil && err == nil {
		err = fmt.Errorf("failed to resolve request %s %s", req.Name, req.Type.ID())
	}
	return
}

func (c *Context) compile(p provider.Provider) (res node.Node, err error) {
	if p == nil {
		return nil, nil
	}

	base, err := c.compile(p.Parent())
	if err != nil {
		return
	}

	deps := make([]node.Node, len(p.Dependencies()))
	for i, dep := range p.Dependencies() {
		deps[i], err = c.Resolve(&dep)
		if err != nil {
			return
		}
	}

	res = p.Map(base, deps)
	return
}
