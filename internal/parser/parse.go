package parser

import (
	"errors"
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func Parse(dir string) (pkgCfg *mapper.PackageConfig, err error) {
	pkgCfg = new(mapper.PackageConfig)

	pkg, err := load(dir)
	if err != nil {
		return
	}

	imports := NewImportMap()

	for _, file := range pkg.Syntax {
		mappers := findMappers(file, pkg.TypesInfo)
		if len(mappers) == 0 {
			continue
		}

		imports.Merge(file)

		if pkgCfg.Name != "" {
			if pkgCfg.Name != file.Name.Name {
				err = fmt.Errorf("expected package %s, but got %s", pkgCfg.Name, file.Name.Name)
			}
		}
		pkgCfg.Name = file.Name.Name
		pkgCfg.Mappers = append(pkgCfg.Mappers, mappers...)
	}

	pkgCfg.Imports = imports.ToList()
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
