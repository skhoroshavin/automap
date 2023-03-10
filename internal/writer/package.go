package writer

import (
	"fmt"
	"github.com/skhoroshavin/automap/internal/mapper/ast"
	"io"
)

func WritePackage(out io.Writer, pkg *ast.Package) (err error) {
	_, err = fmt.Fprintf(out, packageHeader, pkg.Name)
	if err != nil {
		return
	}

	err = writeImports(out, pkg.Imports)
	if err != nil {
		return
	}

	for _, mapper := range pkg.Mappers {
		err = writeMapper(out, mapper)
		if err != nil {
			return
		}
	}
	return err
}

const packageHeader = `// Code generated by automap. DO NOT EDIT.

//go:build !automap

//go:generate automap

package %s
`

func writeImports(out io.Writer, imports []string) (err error) {
	if len(imports) == 0 {
		return
	}

	_, err = fmt.Fprintln(out, "\nimport (")
	if err != nil {
		return
	}

	for _, s := range imports {
		_, err = fmt.Fprintf(out, "\t%s\n", s)
		if err != nil {
			return
		}
	}

	_, err = fmt.Fprintln(out, ")")
	return
}
