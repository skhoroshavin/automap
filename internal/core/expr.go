package core

// Expr represents abstract expression
type Expr interface {
	// isExpr is a marker function for types implementing Expr
	isExpr()
}

func (_ *ValueExpr) isExpr()  {}
func (_ *StructExpr) isExpr() {}

// ValueExpr represents simple value expression
type ValueExpr struct {
	Value string
}

// StructExpr represents structure initialization expression
type StructExpr struct {
	Name      string
	Fields    []FieldExpr
	IsPointer bool
}

// FieldExpr represents structure field initialization expression and can occur only inside StructExpr
type FieldExpr struct {
	Name  string
	Value Expr
}
