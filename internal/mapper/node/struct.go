package node

import "github.com/skhoroshavin/automap/internal/mapper/ast"

func NewStruct(name string, fields ...*Field) *Struct {
	return &Struct{
		name:   name,
		fields: fields,
	}
}

func NewStructPtr(name string, fields ...*Field) *Struct {
	return &Struct{
		name:      name,
		fields:    fields,
		isPointer: true,
	}
}

func NewField(name string, value Node) *Field {
	return &Field{
		name:  name,
		value: value,
	}
}

type Struct struct {
	name      string
	fields    []*Field
	isPointer bool
}

type Field struct {
	name  string
	value Node
}

func (s *Struct) Build(mapper *ast.Mapper) ast.Expr {
	res := &ast.StructExpr{
		Name:      s.name,
		IsPointer: s.isPointer,
		Fields:    make([]*ast.Field, len(s.fields)),
	}

	for i, field := range s.fields {
		res.Fields[i] = ast.NewField(field.name, field.value.Build(mapper))
	}

	return res
}
