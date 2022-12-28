package mapper

import "strings"

// Type represents generic type
type Type struct {
	// Name is a type name
	Name string
	// Scope is a type scope (empty for built-in types, or package ID for others)
	Scope string
	// IsPointer marks whether type is a pointer
	IsPointer bool
	// Methods is a list of methods
	Methods []Func
	// IsStruct marks whether type is a structure with all fields public
	IsStruct bool
	// Fields is a list of structure fields
	Fields []Var
}

// ID returns unique type identifier which can be used for comparison
func (t *Type) ID() string {
	s := strings.Builder{}
	if t.IsPointer {
		s.WriteString("*")
	}
	if t.Scope != "" {
		s.WriteString(t.Scope)
		s.WriteString(".")
	}
	s.WriteString(t.Name)
	return s.String()
}

// Var represents name-type pair, which can be variable, field or argument
type Var struct {
	Name string
	Type *Type
}

// Func represents function or method signature
type Func struct {
	Name    string
	Args    []Var
	Result  *Type
	CanFail bool
}
