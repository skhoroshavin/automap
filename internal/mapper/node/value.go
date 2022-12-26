package node

import "github.com/skhoroshavin/automap/internal/mapper/ast"

func NewValue(v string) *Value {
	return &Value{v: v}
}

type Value struct {
	v string
}

func (v *Value) Build(_ *ast.Mapper) ast.Expr {
	return &ast.ValueExpr{Value: v.v}
}
