package internal

import (
	"github.com/skhoroshavin/automap/internal/mapper"
	"github.com/skhoroshavin/automap/internal/parser"
	"github.com/skhoroshavin/automap/internal/writer"
	"io"
)

func AutoMap(out io.Writer, dir string) error {
	parseRes, err := parser.Parse(dir)
	if err != nil {
		return err
	}

	mappings, err := mapper.New(parseRes)
	if err != nil {
		return nil
	}

	return writer.Write(out, mappings)
}
