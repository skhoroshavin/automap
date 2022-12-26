package oldmapper

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper"
	"github.com/skhoroshavin/automap/internal/mapper/ast"
)

func BuildPackage(reg *Registry) (*ast.Package, error) {
	pkg := &ast.Package{
		Name:    reg.Package(),
		Imports: reg.Imports(),
		Mappers: make([]*ast.Mapper, len(reg.Mappers())),
	}

	for i, oldMapper := range reg.Mappers() {
		node := oldMapper.ToType.BuildMapper(mapper.ProviderList{
			{Name: oldMapper.FromName, Type: oldMapper.FromType},
		})
		if node == nil {
			return nil, fmt.Errorf("failed to map from %s to %s", oldMapper.FromType.Name(), oldMapper.ToType.Name())
		}

		mpr := &ast.Mapper{
			Signature: fmt.Sprintf(
				"func %s(%s %s%s) %s%s",
				oldMapper.Name,
				oldMapper.FromName,
				ptr("*", oldMapper.FromType.IsPointer()),
				oldMapper.FromType.Name(),
				ptr("*", oldMapper.ToType.IsPointer()),
				oldMapper.ToType.Name(),
			),
		}
		mpr.Result = node.Build(mpr)

		pkg.Mappers[i] = mpr
	}

	return pkg, nil
}

func ptr(prefix string, isPointer bool) string {
	if isPointer {
		return prefix
	} else {
		return ""
	}
}
