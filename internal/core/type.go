package core

import (
	"fmt"
	"strings"
)

type Provider struct {
	Name string
	Type Type
}

type ProviderList []Provider

func (l ProviderList) FindAccessor(name string, typeName string) string {
	for _, p := range l {
		if p.Type.Name() == typeName {
			if name == "" {
				return p.Name
			}
			if strings.ToLower(name) == strings.ToLower(p.Name) {
				return p.Name
			}
			continue
		}

		sub := p.Type.FindAccessor(name, typeName)
		if sub != "" {
			return fmt.Sprintf("%s.%s", p.Name, sub)
		}
	}
	return ""
}

type Type interface {
	Name() string
	IsPointer() bool // TODO: Remove
	FindAccessor(name string, typeName string) string
	BuildMapper(args ProviderList) Node
}

type OpaqueType struct {
	Name_ string
}

func (t *OpaqueType) Name() string {
	return t.Name_
}

func (t *OpaqueType) IsPointer() bool {
	return false
}

func (t *OpaqueType) FindAccessor(name string, typeName string) string {
	return ""
}

func (t *OpaqueType) BuildMapper(args ProviderList) Node {
	accessor := args.FindAccessor("", t.Name_)
	if accessor != "" {
		return &ValueNode{Value: accessor}
	}

	return nil
}

type StructType struct {
	Name_      string
	IsPointer_ bool
	Fields     ProviderList
	Getters    ProviderList
}

func (t *StructType) Name() string {
	return t.Name_
}

func (t *StructType) IsPointer() bool {
	return t.IsPointer_
}

func (t *StructType) FindAccessor(name string, typeName string) string {
	res := t.Fields.FindAccessor(name, typeName)
	if res != "" {
		return res
	}

	res = t.Getters.FindAccessor(name, typeName)
	if res != "" {
		return fmt.Sprintf("%s()", res)
	}

	return ""
}

func (t *StructType) BuildMapper(args ProviderList) Node {
	accessor := args.FindAccessor("", t.Name_)
	if accessor != "" {
		return &ValueNode{Value: accessor}
	}

	res := &StructNode{
		Name:      t.Name(),
		IsPointer: t.IsPointer(),
		Fields:    make([]NamedNode, len(t.Fields)),
	}
	for i, v := range t.Fields {
		accessor := args.FindAccessor(v.Name, v.Type.Name())
		if accessor == "" {
			return nil
		}
		res.Fields[i].Name = v.Name
		res.Fields[i].Value = &ValueNode{Value: accessor}
	}

	return res
}
