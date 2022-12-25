package internal

import (
	"automap/internal/mapper"
	"automap/internal/parser"
	"automap/internal/writer"
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
