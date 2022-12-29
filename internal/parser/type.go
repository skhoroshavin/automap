package parser

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/types"
	gotypes "go/types"
	"strings"
)

func parseType(t gotypes.Type, pkg *Package, imports Imports) (res *types.Type, err error) {
	res = new(types.Type)

	if ptr, ok := t.(*gotypes.Pointer); ok {
		res.IsPointer = true
		t = ptr.Elem()
	}

	// TODO: Properly parse scope
	res.Name, err = parseTypeName(t.String(), pkg, imports)
	if err != nil {
		return
	}

	var ok bool
	namedType, ok := t.(*gotypes.Named)
	if !ok {
		return
	}
	for i := 0; i != namedType.NumMethods(); i++ {
		gomethod := namedType.Method(i)
		sig, ok := gomethod.Type().(*gotypes.Signature)
		if !ok {
			continue
		}

		if sig.Params().Len() != 0 {
			continue
		}
		if sig.Results().Len() != 1 {
			continue
		}

		method := types.Func{
			Name: gomethod.Name(),
		}
		method.ReturnType, err = parseType(sig.Results().At(0).Type(), pkg, imports)
		if err != nil {
			return
		}

		res.Methods = append(res.Methods, method)
	}

	structType, ok := namedType.Underlying().(*gotypes.Struct)
	if !ok {
		return
	}
	res.IsStruct = true
	for i := 0; i != structType.NumFields(); i++ {
		gofield := structType.Field(i)
		if !gofield.Exported() {
			res.IsStruct = false
			continue
		}

		field := types.Var{
			Name: gofield.Name(),
		}
		field.Type, err = parseType(gofield.Type(), pkg, imports)
		if err != nil {
			return
		}

		res.Fields = append(res.Fields, field)
	}

	return res, nil
}

func parseTypeName(typeId string, pkg *Package, imports Imports) (string, error) {
	name := typeId

	dot := strings.LastIndex(name, ".")
	if dot < 0 {
		return name, nil
	}

	path := name[:dot]
	name = name[dot+1:]
	if path == pkg.Path {
		return name, nil
	}

	prefix, ok := imports[path]
	if !ok {
		return "", fmt.Errorf("failed to resolve type name %s", typeId)
	}

	return fmt.Sprintf("%s.%s", prefix, name), nil
}
