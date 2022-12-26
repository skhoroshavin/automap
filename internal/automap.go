package internal

import (
	"github.com/skhoroshavin/automap/internal/mapper"
	"github.com/skhoroshavin/automap/internal/parser"
	"github.com/skhoroshavin/automap/internal/writer"
	"io"
)

func AutoMap(out io.Writer, dir string) error {
	pkgConfig, err := parser.Parse(dir)
	if err != nil {
		return err
	}

	pkg, err := mapper.BuildPackage(pkgConfig)
	if err != nil {
		return err
	}

	return writer.WritePackage(out, pkg)
}
