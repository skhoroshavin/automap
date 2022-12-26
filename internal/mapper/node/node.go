package node

import "github.com/skhoroshavin/automap/internal/mapper/ast"

type Node interface {
	Build(*ast.Mapper) ast.Expr
}

// TODO: Either get rid of or move into testing
func build(node Node) *ast.Mapper {
	res := new(ast.Mapper)
	res.Result = node.Build(res)
	return res
}
