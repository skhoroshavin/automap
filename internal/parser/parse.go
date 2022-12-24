package parser

import (
	"errors"
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

type Result struct {
	Package  string
	Imports  Imports
	TypeInfo *types.Info
	Mappers  []*Mapper
}

func Parse(dir string) (res *Result, err error) {
	res = new(Result)

	pkg, err := load(dir)
	if err != nil {
		return
	}

	res.TypeInfo = pkg.TypesInfo
	res.Imports = newImports()

	for _, file := range pkg.Syntax {
		mappers := findMappers(file)
		if len(mappers) == 0 {
			continue
		}

		if res.Package != "" {
			if res.Package != file.Name.Name {
				err = fmt.Errorf("expected package %s, but got %s", res.Package, file.Name.Name)
			}
		}
		res.Package = file.Name.Name

		mergeImports(res.Imports, file)

		res.Mappers = append(res.Mappers, mappers...)
	}

	return
}

func load(dir string) (res *packages.Package, err error) {
	cfg := &packages.Config{
		Mode:       packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
		Dir:        dir,
		BuildFlags: []string{"-tags", "automap"},
	}

	pkgs, err := packages.Load(cfg)
	if err != nil {
		return
	}

	if len(pkgs) != 1 {
		err = errors.New("loaded more than one package")
		return
	}

	res = pkgs[0]
	if len(res.Errors) > 0 {
		err = fmt.Errorf("syntax errors: %+v", res.Errors)
		return
	}

	return
}

func findMappers(file *ast.File) (res []*Mapper) {
	ast.Inspect(file, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.File:
			return true
		case *ast.FuncDecl:
			mapper := parseMapper(x)
			if mapper != nil {
				res = append(res, mapper)
			}
			return false
		default:
			return false
		}
	})
	return
}
