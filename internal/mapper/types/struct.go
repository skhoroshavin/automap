package types

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/node"
)

func NewStruct(name string, fields ProviderList, getters ProviderList) *Struct {
	return &Struct{
		name:    name,
		fields:  fields,
		getters: getters,
	}
}

func NewStructPtr(name string, fields ProviderList, getters ProviderList) *Struct {
	return &Struct{
		name:      name,
		isPointer: true,
		fields:    fields,
		getters:   getters,
	}
}

type Struct struct {
	name      string
	isPointer bool
	fields    ProviderList
	getters   ProviderList
}

func (t *Struct) Name() string {
	return t.name
}

func (t *Struct) IsPointer() bool {
	return t.isPointer
}

func (t *Struct) FindAccessor(name string, typeName string) string {
	res := t.fields.FindAccessor(name, typeName, false)
	if res != "" {
		return res
	}

	res = t.getters.FindAccessor(name, typeName, true)
	if res != "" {
		return res
	}

	return ""
}

func (t *Struct) BuildMapper(args ProviderList) (node.Node, error) {
	accessor := args.FindAccessor("", t.name, false)
	if accessor != "" {
		return node.NewValue(accessor), nil
	}

	fields := make([]*node.Field, len(t.fields))
	for i, v := range t.fields {
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
