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
