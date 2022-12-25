package core

// FuncBody represents function body
type FuncBody struct {
	// Vars is a list of variable assignments
	Vars []Variable
	// Result is a final return statement
	Result Expr
}

// Variable represents local variable in a function
type Variable struct {
	Name  string
	Value Expr
}

// Expr represents abstract expression
type Expr interface {
	// isExpr is a marker function for types implementing Expr
	isExpr()
}

// ValueExpr represents simple value expression
type ValueExpr struct {
	Value string
}

func (_ *ValueExpr) isExpr() {}

// StructExpr represents structure initialization expression
type StructExpr struct {
	Name      string
	Fields    []FieldExpr
	IsPointer bool
}

func (_ *StructExpr) isExpr() {}

// FieldExpr represents structure field initialization expression and can occur only inside StructExpr
type FieldExpr struct {
	Name  string
	Value Expr
}
