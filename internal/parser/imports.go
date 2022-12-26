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

func (m Imports) Merge(file *ast.File) {
	for _, spec := range file.Imports {
		m.Insert(spec)
	}
}

func (m Imports) Insert(spec *ast.ImportSpec) {
	path := strings.ReplaceAll(spec.Path.Value, "\"", "")

	var name string
	if spec.Name == nil {
		idx := strings.LastIndex(path, "/")
		if idx < 0 {
			name = path
		} else {
			name = path[idx+1:]
		}
	} else {
		name = spec.Name.Name
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
