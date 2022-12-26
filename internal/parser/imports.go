package parser

import (
	"fmt"
	"go/ast"
	"strings"
)

type ImportMap map[string]string

func NewImportMap() ImportMap {
	return make(ImportMap)
}

func (m ImportMap) Merge(file *ast.File) {
	for _, spec := range file.Imports {
		path, name := parseImport(spec)
		if name == "automap" {
			continue
		}
		m[path] = name
	}
}

func (m ImportMap) ToList() []string {
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

func parseImport(spec *ast.ImportSpec) (path, name string) {
	path = strings.ReplaceAll(spec.Path.Value, "\"", "")
	if spec.Name != nil {
		name = spec.Name.Name
		return
	}
	idx := strings.LastIndex(path, "/")
	if idx < 0 {
		name = path
		return
	}
	name = path[idx+1:]
	return
}
