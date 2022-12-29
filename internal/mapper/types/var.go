package types

// Var represents name-type pair, which can be variable, field or argument
type Var struct {
	Name string
	Type *Type
}
