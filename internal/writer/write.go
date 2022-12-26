package writer

import (
	"github.com/skhoroshavin/automap/internal/mapper"
	"io"
)

func Write(out io.Writer, reg *mapper.Registry) (err error) {
	err = writeHeader(out, reg.Package())
	if err != nil {
		return
	}

	err = writeImports(out, reg.Imports())
	if err != nil {
		return
	}

	for _, mapper := range reg.Mappers() {
		err = writeOldMapper(out, mapper)
		if err != nil {
			return
		}
	}

	return
}
