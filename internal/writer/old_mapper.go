package writer

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/core"
	"github.com/skhoroshavin/automap/internal/mapper"
	"io"
)

func writeOldMapper(out io.Writer, mapper *mapper.Mapper) (err error) {
	_, err = fmt.Fprintf(
		out,
		"\nfunc %s(%s %s%s) %s%s {\n",
		mapper.Name,
		mapper.FromName,
		ptr("*", mapper.FromType.IsPointer()),
		mapper.FromType.Name(),
		ptr("*", mapper.ToType.IsPointer()),
		mapper.ToType.Name(),
	)
	if err != nil {
		return
	}

	node := mapper.ToType.BuildMapper(core.ProviderList{
		{Name: mapper.FromName, Type: mapper.FromType},
	})
	if node == nil {
		err = fmt.Errorf("failed to map from %s to %s", mapper.FromType.Name(), mapper.ToType.Name())
		return
	}

	err = writeMapper(out, core.BuildFuncBody(node))
	if err != nil {
		return
	}

	_, err = fmt.Fprintln(out, "}")
	return
}

func ptr(prefix string, isPointer bool) string {
	if isPointer {
		return prefix
	} else {
		return ""
	}
}
