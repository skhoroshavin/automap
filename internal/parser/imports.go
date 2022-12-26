package parser

import (
	"fmt"
	"go/ast"
	"strings"
)

type Imports map[string]string

func NewImports() Imports {
	return make(Imports)
}

func (m Imports) ParseFile(file *ast.File) {
	for _, spec := range file.Imports {
		m.ParseImport(spec)
	}
}

func (m Imports) ParseImport(spec *ast.ImportSpec) {
	path := strings.ReplaceAll(spec.Path.Value, "\"", "")

	var name string
	if spec.Name != nil {
		name = spec.Name.Name
	} else {
		name = parsePackageName(path)
	}

	if name == "automap" {
		return
	}

	m[path] = name
}

func (m Imports) ToList() []string {
	res := make([]string, 0, len(m))
	for path, name := range m {
		isUnnamed := (path == name) ||
			strings.HasSuffix(path, fmt.Sprintf("/%s", name))

		if isUnnamed {
			res = append(res, fmt.Sprintf(`"%s"`, path))
		} else {
			res = append(res, fmt.Sprintf(`%s "%s"`, name, path))
		}
	}
	return res
}
