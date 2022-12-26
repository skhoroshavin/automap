package writer

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/core"
	"github.com/skhoroshavin/automap/internal/core/ast"
	"github.com/skhoroshavin/automap/internal/mapper"
	"io"
)

func writeOldMapper(out io.Writer, oldMapper *mapper.Mapper) error {
	node := oldMapper.ToType.BuildMapper(core.ProviderList{
		{Name: oldMapper.FromName, Type: oldMapper.FromType},
	})
	if node == nil {
		return fmt.Errorf("failed to map from %s to %s", oldMapper.FromType.Name(), oldMapper.ToType.Name())
	}

	mapper := &ast.Mapper{
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
	mapper.Result = node.Build(mapper)

	return writeMapper(out, mapper)
}

func ptr(prefix string, isPointer bool) string {
	if isPointer {
		return prefix
	} else {
		return ""
	}
}
