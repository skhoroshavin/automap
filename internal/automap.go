package internal

import (
	"github.com/skhoroshavin/automap/internal/oldmapper"
	"github.com/skhoroshavin/automap/internal/parser"
	"github.com/skhoroshavin/automap/internal/writer"
	"io"
)

func AutoMap(out io.Writer, dir string) error {
	parseRes, err := parser.Parse(dir)
	if err != nil {
		return err
	}

	oldReg, err := oldmapper.New(parseRes)
	if err != nil {
		return nil
	}

	pkg, err := oldmapper.BuildPackage(oldReg)
	if err != nil {
		return err
	}

	return writer.WritePackage(out, pkg)
}
