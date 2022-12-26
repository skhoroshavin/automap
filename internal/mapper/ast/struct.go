package ast

// NewStruct creates structure initialization expression
func NewStruct(name string, fields ...Field) *StructExpr {
	return &StructExpr{
		Name:   name,
		Fields: fields,
	}
}

// NewStructPtr creates pointer to structure initialization expression
func NewStructPtr(name string, fields ...Field) *StructExpr {
	return &StructExpr{
		Name:      name,
		Fields:    fields,
		IsPointer: true,
	}
}

// NewField creates new field
func NewField(name string, value Expr) Field {
	return Field{
		Name:  name,
		Value: value,
	}
}

// StructExpr represents structure initialization expression
type StructExpr struct {
	Name      string
	Fields    []Field
	IsPointer bool
}

func (_ *StructExpr) isExpr() {}

// Field represents structure field initialization
type Field struct {
	Name  string
	Value Expr
}
