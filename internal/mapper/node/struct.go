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

func (s *Struct) CompileTo(mapper *ast.Mapper) {
	fields := make([]*ast.Field, len(s.fields))
	for i, field := range s.fields {
		field.value.CompileTo(mapper)
		fields[i] = ast.NewField(field.name, mapper.Result)
	}

	if s.isPointer {
		mapper.Result = ast.NewStructPtr(s.name, fields...)
	} else {
		mapper.Result = ast.NewStruct(s.name, fields...)
	}
}
