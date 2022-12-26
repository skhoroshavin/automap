package mapper

import "github.com/skhoroshavin/automap/internal/mapper/ast"

type Node interface {
	Build(*ast.Mapper) ast.Expr
}

func BuildFuncBody(node Node) *ast.Mapper {
	res := new(ast.Mapper)
	res.Result = node.Build(res)
	return res
}

type NamedNode struct {
	Name  string
	Value Node
}

type ValueNode struct {
	Value string
}

func (n *ValueNode) Build(mapper *ast.Mapper) ast.Expr {
	return &ast.ValueExpr{Value: n.Value}
}

type StructNode struct {
	Name      string
	Fields    []NamedNode
	IsPointer bool
}

func (n *StructNode) Build(mapper *ast.Mapper) ast.Expr {
	res := &ast.StructExpr{
		Name:      n.Name,
		IsPointer: n.IsPointer,
		Fields:    make([]*ast.Field, len(n.Fields)),
	}

	for i, field := range n.Fields {
		res.Fields[i] = ast.NewField(field.Name, field.Value.Build(mapper))
	}

	return res
}

type FuncNode struct {
	Name string

	FuncName string
	Inputs   []Node
}
