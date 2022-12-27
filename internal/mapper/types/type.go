package types

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/node"
	"strings"
)

type Type interface {
	Name() string
	IsPointer() bool // TODO: Remove
	FindAccessor(name string, typeName string) string
	BuildMapper(args ProviderList) (node.Node, error)
}

type Provider struct {
	Name string
	Type Type
}

type ProviderList []Provider

func (l ProviderList) FindAccessor(name string, typeName string, isGetter bool) string {
	for _, p := range l {
		accessor := p.Name
		if isGetter {
			accessor = fmt.Sprintf("%s()", accessor)
		}

		if p.Type.Name() == typeName {
			if name == "" {
				return accessor
			}
			if strings.ToLower(name) == strings.ToLower(p.Name) {
				return accessor
			}
			continue
		}

		if strings.HasPrefix(strings.ToLower(name), strings.ToLower(p.Name)) {
			name := name[len(p.Name):]
			sub := p.Type.FindAccessor(name, typeName)
			if sub != "" {
				return fmt.Sprintf("%s.%s", accessor, sub)
			}
		}

		sub := p.Type.FindAccessor(name, typeName)
		if sub != "" {
			return fmt.Sprintf("%s.%s", accessor, sub)
		}
	}
	return ""
}
