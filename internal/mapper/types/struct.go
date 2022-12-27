package types

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/node"
)

type Struct struct {
	Name_      string
	IsPointer_ bool
	Fields     ProviderList
	Getters    ProviderList
}

func (t *Struct) Name() string {
	return t.Name_
}

func (t *Struct) IsPointer() bool {
	return t.IsPointer_
}

func (t *Struct) FindAccessor(name string, typeName string) string {
	res := t.Fields.FindAccessor(name, typeName, false)
	if res != "" {
		return res
	}

	res = t.Getters.FindAccessor(name, typeName, true)
	if res != "" {
		return res
	}

	return ""
}

func (t *Struct) BuildMapper(args ProviderList) (node.Node, error) {
	accessor := args.FindAccessor("", t.Name_, false)
	if accessor != "" {
		return node.NewValue(accessor), nil
	}

	fields := make([]*node.Field, len(t.Fields))
	for i, v := range t.Fields {
		accessor := args.FindAccessor(v.Name, v.Type.Name(), false)
		if accessor == "" {
			return nil, fmt.Errorf("no accessor found for field %s %s", v.Name, v.Type.Name())
		}
		fields[i] = node.NewField(v.Name, node.NewValue(accessor))
	}

	if t.IsPointer() {
		return node.NewStructPtr(t.Name(), fields...), nil
	} else {
		return node.NewStruct(t.Name(), fields...), nil
	}
}
