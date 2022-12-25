package writer

import (
	"automap/internal/mapper"
	"fmt"
	"io"
)

func writeMapper(out io.Writer, mapper *mapper.Mapper) (err error) {
	_, err = fmt.Fprintf(
		out,
		"\nfunc %s(%s %s%s) %s%s {\n",
		mapper.Name,
		mapper.FromName,
		ptr("*", mapper.FromType.IsPointer),
		mapper.FromType.Name,
		ptr("*", mapper.ToType.IsPointer),
		mapper.ToType.Name,
	)
	if err != nil {
		return
	}

	_, err = fmt.Fprintf(
		out,
		"\treturn %s%s{\n",
		ptr("&", mapper.ToType.IsPointer),
		mapper.ToType.Name,
	)
	if err != nil {
		return
	}

	for i := 0; i != mapper.ToType.Struct.NumFields(); i++ {
		toField := mapper.ToType.Struct.Field(i)
		accessor := mapper.FromType.FindAccessor(toField.Name(), toField.Type())
		if accessor == "" {
			return fmt.Errorf("cannot map %s", toField.String())
		}
		_, err = fmt.Fprintf(out, "\t\t%s: %s.%s,\n", toField.Name(), mapper.FromName, accessor)
		if err != nil {
			return
		}
	}

	_, err = fmt.Fprintln(out, "\t}\n}")
	return
}

func ptr(prefix string, isPointer bool) string {
	if isPointer {
		return prefix
	} else {
		return ""
	}
}
