package parser

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/types"
	gotypes "go/types"
	"strings"
)

func parseType(t gotypes.Type, pkg *Package, imports Imports) (types.Type, error) {
	isPointer := false
	if ptr, ok := t.(*gotypes.Pointer); ok {
		isPointer = true
		t = ptr.Elem()
	}

	name, err := parseTypeName(t.String(), pkg, imports)
	if err != nil {
		return nil, err
	}

	var ok bool
	namedType, ok := t.(*gotypes.Named)
	if !ok {
		return types.NewOpaque(name), nil
	}

	structType, ok := namedType.Underlying().(*gotypes.Struct)
	if !ok {
		return types.NewOpaque(name), nil
	}

	fields := make(types.ProviderList, 0, structType.NumFields())
	for i := 0; i != structType.NumFields(); i++ {
		field := structType.Field(i)
		if !field.Exported() {
			continue
		}
		typ, err := parseType(field.Type(), pkg, imports)
		if err != nil {
			return nil, err
		}

		fields = append(fields, types.Provider{
			Name: field.Name(),
			Type: typ,
		})
	}

	getters := make(types.ProviderList, 0, namedType.NumMethods())
	for i := 0; i != namedType.NumMethods(); i++ {
		method := namedType.Method(i)
		sig, ok := method.Type().(*gotypes.Signature)
		if !ok {
			continue
		}
		if sig.Results().Len() != 1 {
			continue
		}
		if sig.Params().Len() != 0 {
			continue
		}
		typ, err := parseType(sig.Results().At(0).Type(), pkg, imports)
		if err != nil {
			return nil, err
		}

		getters = append(getters, types.Provider{
			Name: method.Name(),
			Type: typ,
		})
	}

	if isPointer {
		return types.NewStructPtr(name, fields, getters), nil
	} else {
		return types.NewStruct(name, fields, getters), nil
	}
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
