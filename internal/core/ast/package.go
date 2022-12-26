package ast

// Package represents package
type Package struct {
	Name    string
	Imports []string
	Mappers []*Mapper
}

// Mapper represents mapper
type Mapper struct {
	// Signature is a mapper signature (TODO: Make it an object)
	Signature string
	// Vars is a list of variable assignments
	Vars []Variable
	// Result is a final return statement
	Result Expr
}

// Variable represents local variable in a mapper
type Variable struct {
	Name  string
	Value Expr
}

// Expr represents abstract expression
type Expr interface {
	// isExpr is a marker function for types implementing Expr
	isExpr()
}
