package mapper

import "github.com/skhoroshavin/automap/internal/mapper/ast"

// PackageConfig is a config for building package with mappers
type PackageConfig struct {
	Name    string
	Imports []string
	Mappers []*Config
}

// BuildPackage builds the package with mappers
func BuildPackage(cfg *PackageConfig) (*ast.Package, error) {
	pkg := &ast.Package{
		Name:    cfg.Name,
		Imports: cfg.Imports,
		Mappers: make([]*ast.Mapper, len(cfg.Mappers)),
	}

	for i, mapperCfg := range cfg.Mappers {
		mapper, err := Build(mapperCfg)
		if err != nil {
			return nil, err
		}
		pkg.Mappers[i] = mapper
	}

	return pkg, nil
}
