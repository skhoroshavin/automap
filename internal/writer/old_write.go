package writer

import (
	"github.com/skhoroshavin/automap/internal/oldmapper"
	"io"
)

func OldWrite(out io.Writer, reg *oldmapper.Registry) (err error) {
	err = writeOldHeader(out, reg.Package())
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
