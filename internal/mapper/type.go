package mapper

import (
	"automap/internal/utils"
	"fmt"
	"go/ast"
	"go/types"
)

func NewType(typeExpr ast.Expr, typeInfo *types.Info) (res *Type, err error) {
	res = new(Type)

	// Dereference pointer if needed
	if starExpr, ok := typeExpr.(*ast.StarExpr); ok {
		res.IsPointer = true
		typeExpr = starExpr.X
	}

	// Name
	res.Name = utils.AST2String(typeExpr)

	// Get type
	typ := typeInfo.TypeOf(typeExpr)
	if typ == nil {
		err = fmt.Errorf("unknown type %s", res.Name)
		return
	}

	var ok bool
	res.Named, ok = typ.(*types.Named)
	if !ok {
		err = fmt.Errorf("type %s is not named", res.Name)
		return
	}

	res.Struct, ok = res.Named.Underlying().(*types.Struct)
	if !ok {
		err = fmt.Errorf("type %s is not struct", res.Name)
		return
	}

	return
}

type Type struct {
	Name      string
	Named     *types.Named
	Struct    *types.Struct
	IsPointer bool
}

func (t *Type) FindAccessor(name string, typ types.Type) string {
	for i := 0; i != t.Struct.NumFields(); i++ {
		f := t.Struct.Field(i)
		if f.Name() != name {
			continue
		}
		if f.Type() != typ {
			continue
		}
		return f.Name()
	}

	return ""
}
