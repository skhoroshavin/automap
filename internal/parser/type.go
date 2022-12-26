package parser

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper"
	"github.com/skhoroshavin/automap/internal/utils"
	"go/ast"
	"go/types"
)

func ParseType(typeExpr ast.Expr, typeInfo *types.Info) (mapper.Type, error) {
	isPointer := false

	// Dereference pointer if needed
	if starExpr, ok := typeExpr.(*ast.StarExpr); ok {
		isPointer = true
		typeExpr = starExpr.X
	}

	// Name
	name := utils.AST2String(typeExpr)

	// Get type
	typ := typeInfo.TypeOf(typeExpr)
	if typ == nil {
		return nil, fmt.Errorf("unknown type %s", name)
	}

	var ok bool
	namedType, ok := typ.(*types.Named)
	if !ok {
		return &mapper.OpaqueType{Name_: name}, nil
	}

	structType, ok := namedType.Underlying().(*types.Struct)
	if !ok {
		return &mapper.OpaqueType{Name_: name}, nil
	}

	res := &mapper.StructType{
		Name_:      name,
		IsPointer_: isPointer,
		Fields:     make(mapper.ProviderList, 0, structType.NumFields()),
		Getters:    make(mapper.ProviderList, 0, namedType.NumMethods()),
	}
	for i := 0; i != structType.NumFields(); i++ {
		field := structType.Field(i)
		if !field.Exported() {
			continue
		}
		res.Fields = append(res.Fields, mapper.Provider{
			Name: field.Name(),
			Type: &mapper.OpaqueType{Name_: field.Type().String()},
		})
	}
	for i := 0; i != namedType.NumMethods(); i++ {
		method := namedType.Method(i)
		sig, ok := method.Type().(*types.Signature)
		if !ok {
			continue
		}
		if sig.Results().Len() != 1 {
			continue
		}
		if sig.Params().Len() != 0 {
			continue
		}
		res.Getters = append(res.Getters, mapper.Provider{
			Name: method.Name(),
			Type: &mapper.OpaqueType{Name_: sig.Results().At(0).Type().String()},
		})
	}

	return res, nil
}
