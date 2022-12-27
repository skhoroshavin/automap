package types

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/node"
)

type Opaque struct {
	Name_ string
}

func (t *Opaque) Name() string {
	return t.Name_
}

func (t *Opaque) IsPointer() bool {
	return false
}

func (t *Opaque) FindAccessor(name string, typeName string) string {
	return ""
}

func (t *Opaque) BuildMapper(args ProviderList) (node.Node, error) {
	accessor := args.FindAccessor("", t.Name_, false)
	if accessor != "" {
		return node.NewValue(accessor), nil
	}

	return nil, fmt.Errorf("no accessor found for type %s", t.Name_)
}
