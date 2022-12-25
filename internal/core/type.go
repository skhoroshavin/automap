package core

import "strings"

type Provider struct {
	Name string
	// TODO: Make it type
	Type string
}

type StructType struct {
	Name      string
	IsPointer bool
	Fields    []Provider
	Getters   []Provider
}

func (t *StructType) BuildMapper(args []Provider) Node {
	for _, arg := range args {
		if arg.Type == t.Name {
			return &ValueNode{Value: arg.Name}
		}
	}

	res := &StructNode{
		Name:      t.Name,
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
