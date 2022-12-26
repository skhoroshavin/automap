package core

import "strings"

type Provider struct {
	Name string
	Type Type
}

type Type interface {
	Name() string
}

type OpaqueType struct {
	Name_ string
}

func (t *OpaqueType) Name() string {
	return t.Name_
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

func (t *StructType) BuildMapper(args []Provider) Node {
	for _, arg := range args {
		if arg.Type.Name() == t.Name() {
			return &ValueNode{Value: arg.Name}
		}
	}

	res := &StructNode{
		Name:      t.Name(),
		IsPointer: t.IsPointer,
		Fields:    make([]NamedNode, len(t.Fields)),
	}
	for i, v := range t.Fields {
		for _, arg := range args {
			if strings.ToLower(arg.Name) != strings.ToLower(v.Name) {
				continue
			}
			if arg.Type != v.Type {
				continue
			}
			res.Fields[i].Name = v.Name
			res.Fields[i].Value = &ValueNode{Value: arg.Name}
			break
		}
		if res.Fields[i].Value == nil {
			return nil
		}
	}

	return res
}
