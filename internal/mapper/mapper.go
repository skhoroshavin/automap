package mapper

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/ast"
	"github.com/skhoroshavin/automap/internal/mapper/types"
)

// TODO: Make more generic
type Config struct {
	Name     string
	FromName string
	FromType types.Type
	ToType   types.Type
}

func Build(cfg *Config) (*ast.Mapper, error) {
	node, err := cfg.ToType.BuildMapper(types.ProviderList{
		{Name: cfg.FromName, Type: cfg.FromType},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to map from %s to %s: %w", cfg.FromType.Name(), cfg.ToType.Name(), err)
	}

	mapper := &ast.Mapper{
		Signature: fmt.Sprintf(
			"func %s(%s %s%s) %s%s",
			cfg.Name,
			cfg.FromName,
			ptr("*", cfg.FromType.IsPointer()),
			cfg.FromType.Name(),
			ptr("*", cfg.ToType.IsPointer()),
			cfg.ToType.Name(),
		),
	}
	node.CompileTo(mapper)

	return mapper, nil
}

func ptr(prefix string, isPointer bool) string {
	if isPointer {
		return prefix
	} else {
		return ""
	}
}
