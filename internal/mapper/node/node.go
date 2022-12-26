package node

import "github.com/skhoroshavin/automap/internal/mapper/ast"

type Node interface {
	CompileTo(*ast.Mapper)
}
