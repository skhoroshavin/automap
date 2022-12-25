package parser

import (
	"github.com/skhoroshavin/automap/internal/utils"
	"go/ast"
	"strings"
)

type Imports map[string]struct{}

func newImports() Imports {
	return make(Imports)
}

func mergeImports(imports Imports, file *ast.File) {
	for _, spec := range file.Imports {
		if spec.Path.Value == `"automap"` {
			continue
		}
		if strings.HasSuffix(spec.Path.Value, `/automap"`) {
			continue
		}
		imports[utils.AST2String(spec)] = struct{}{}
	}
}
