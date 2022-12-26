package mapper

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/ast"
)

// TODO: Make more generic
type Config struct {
	Name     string
	FromName string
	FromType Type
	ToType   Type
}

func Build(cfg *Config) (*ast.Mapper, error) {
	node := cfg.ToType.BuildMapper(ProviderList{
		{Name: cfg.FromName, Type: cfg.FromType},
	})
	if node == nil {
		return nil, fmt.Errorf("failed to map from %s to %s", cfg.FromType.Name(), cfg.ToType.Name())
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
	mapper.Result = node.Build(mapper)

	return mapper, nil
}

func ptr(prefix string, isPointer bool) string {
	if isPointer {
		return prefix
	} else {
		return ""
	}
}
