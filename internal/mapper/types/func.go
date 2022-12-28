package types

// Func represents function or method signature
type Func struct {
	Name       string
	Args       []Var
	ReturnType *Type
	CanFail    bool
}
