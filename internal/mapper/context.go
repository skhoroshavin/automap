package mapper

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/node"
)

// Context represents mapping context
type Context struct {
	providers ProviderList
}

// AddProvider adds provider to the list of available providers
func (c *Context) AddProvider(p Provider) {
	c.providers = append(c.providers, p)
}

// Resolve tries to resolve given Request into node tree
func (c *Context) Resolve(req Request) (res node.Node, err error) {
	// TODO: Add implicit resolvers for request type

	c.providers.ForEach(func(p Provider) bool {
		if !p.Match(req) {
			return false
		}

		res, err = c.compile(p)
		if err != nil {
			return false
		}

		return true
	})

	if res == nil && err == nil {
		err = fmt.Errorf("failed to resolve request %s %s", req.Name, req.TypeID)
	}
	return
}

func (c *Context) compile(p Provider) (res node.Node, err error) {
	if p == nil {
		return nil, nil
	}

	base, err := c.compile(p.Parent())
	if err != nil {
		return
	}

	deps := make([]node.Node, len(p.Dependencies()))
	for i, dep := range p.Dependencies() {
		deps[i], err = c.Resolve(dep)
		if err != nil {
			return
		}
	}

	res = p.Map(base, deps)
	return
}
