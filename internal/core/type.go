package core

import "strings"

type Provider struct {
	Name string
	Type Type
}

type ProviderList []Provider

func (l ProviderList) FindAccessor(name string, typeName string) string {
	for _, p := range l {
		if strings.ToLower(p.Name) != strings.ToLower(name) && name != "" {
			continue
		}
		if p.Type.Name() == typeName {
			return p.Name
		}
	}
	return ""
}

type Type interface {
	Name() string
	FindAccessor(name string, typeName string) string
}

type OpaqueType struct {
	Name_ string
}

func (t *OpaqueType) Name() string {
	return t.Name_
}

func (t *OpaqueType) FindAccessor(name string, typeName string) string {
	return ""
}

type StructType struct {
	Name_     string
	IsPointer bool
	Fields    []Provider
	Getters   []Provider
}

func (t *StructType) Name() string {
	return t.Name_
}

func (t *StructType) FindAccessor(name string, typeName string) string {
	return ""
}

func (t *StructType) BuildMapper(args ProviderList) Node {
	accessor := args.FindAccessor("", t.Name_)
	if accessor != "" {
		return &ValueNode{Value: accessor}
	}

	res := &StructNode{
		Name:      t.Name(),
		IsPointer: t.IsPointer,
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
