package ast

// NewValue creates new value expression
func NewValue(v string) *ValueExpr {
	return &ValueExpr{Value: v}
}

// ValueExpr represents simple value expression
type ValueExpr struct {
	Value string
}

func (_ *ValueExpr) isExpr() {}
