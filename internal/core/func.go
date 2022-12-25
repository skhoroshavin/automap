package core

// TODO: Implement expressions so that actual code formatting could be moved to writer

// FuncBody represents function body
type FuncBody struct {
	// Vars is a list of variable assignments
	Vars []Variable
	// Result is a final return statement
	Result string
}

// Variable represents local variable in a function
type Variable struct {
	Name  string
	Value string
}
