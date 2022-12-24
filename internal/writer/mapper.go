package writer

import (
	"automap/internal/mapper"
	"automap/internal/utils"
	"fmt"
	"io"
)

func writeMapper(out io.StringWriter, mapper *mapper.Mapper) (err error) {
	err = utils.WriteLn(
		out,
		"\nfunc %s(%s *%s) *%s {",
		mapper.Name,
		mapper.FromName,
		mapper.FromType.Name,
		mapper.ToType.Name,
	)
	if err != nil {
		return
	}

	err = utils.WriteLn(
		out,
		"\treturn &%s{",
		mapper.ToType.Name)
	if err != nil {
		return
	}

	for i := 0; i != mapper.ToType.Struct.NumFields(); i++ {
		toField := mapper.ToType.Struct.Field(i)
		accessor := mapper.FromType.FindAccessor(toField.Name(), toField.Type())
		if accessor == "" {
			return fmt.Errorf("cannot map field %s", toField.String())
		}
		utils.WriteLn(out, "\t\t%s: %s.%s,", toField.Name(), mapper.FromName, accessor)
	}

	err = utils.WriteLn(out, "\t}\n}")
	return
}
