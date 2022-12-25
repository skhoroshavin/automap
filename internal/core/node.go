package core

type Node interface {
	Build(fn *FuncBody) Expr
}

func BuildFuncBody(node Node) *FuncBody {
	res := new(FuncBody)
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

func (n *ValueNode) Build(_ *FuncBody) Expr {
	return &ValueExpr{Value: n.Value}
}

type StructNode struct {
	Name      string
	Fields    []NamedNode
	IsPointer bool
}

func (n *StructNode) Build(fn *FuncBody) Expr {
	res := &StructExpr{
		Name:      n.Name,
		IsPointer: n.IsPointer,
		Fields:    make([]FieldExpr, len(n.Fields)),
	}

	for i, field := range n.Fields {
		res.Fields[i].Name = field.Name
		res.Fields[i].Value = field.Value.Build(fn)
	}

	return res
}

type FuncNode struct {
	Name string

	FuncName string
	Inputs   []Node
}
