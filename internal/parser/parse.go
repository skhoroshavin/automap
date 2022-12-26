package parser

import (
	"errors"
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func Parse(dir string) (*mapper.PackageConfig, error) {
	gopkg, err := load(dir)
	if err != nil {
		return nil, err
	}

	pkg := ParsePackage(gopkg)
	imports := NewImports()
	var allMappers []*mapper.Config

	for _, file := range gopkg.Syntax {
		mappers := findMappers(file, gopkg.TypesInfo)
		if len(mappers) == 0 {
			continue
		}

		if file.Name.Name != pkg.Name {
			return nil, fmt.Errorf("expected package %s, but got %s", pkg.Name, file.Name.Name)
		}

		imports.ParseFile(file)
		allMappers = append(allMappers, mappers...)
	}

	return &mapper.PackageConfig{
		Name:    pkg.Name,
		Imports: imports.ToList(),
		Mappers: allMappers,
	}, nil
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

func findMappers(file *ast.File, typeInfo *types.Info) (res []*mapper.Config) {
	ast.Inspect(file, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.File:
			return true
		case *ast.FuncDecl:
			parsedMapper := parseMapper(x)
			if parsedMapper == nil {
				return false
			}
			mapperConfig, err := buildMapperConfig(parsedMapper, typeInfo)
			if err != nil {
				return false
			}
			res = append(res, mapperConfig)
			return false
		default:
			return false
		}
	})
	return
}
