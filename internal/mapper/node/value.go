package node

import "github.com/skhoroshavin/automap/internal/mapper/ast"

func NewValue(v string) *Value {
	return &Value{v: v}
}

type Value struct {
	v string
}

func (v *Value) CompileTo(mapper *ast.Mapper) {
	mapper.Result = ast.NewValue(v.v)
}
