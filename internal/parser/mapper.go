package parser

import (
	"go/ast"
)

type Mapper struct {
	Name    string
	From    *Field
	To      *Field
	Builder []ast.Expr
}

type Field struct {
	Name string
	Type ast.Expr
}

func parseMapper(fun *ast.FuncDecl) *Mapper {
	// For now allow only 1:1 mappers, but this may change in future
	if fun.Type.Params.NumFields() != 1 {
		return nil
	}
	if len(fun.Type.Params.List[0].Names) > 1 {
		return nil
	}
	if fun.Type.Results.NumFields() != 1 {
		return nil
	}
	if len(fun.Type.Results.List[0].Names) > 1 {
		return nil
	}

	// Mappers should always include builder expression
	buildExpr := findBuilderExpr(fun)
	if buildExpr == nil {
		return nil
	}

	return &Mapper{
		Name:    fun.Name.Name,
		From:    parseField(fun.Type.Params.List[0]),
		To:      parseField(fun.Type.Results.List[0]),
		Builder: buildExpr.Args,
	}
}

func parseField(field *ast.Field) (res *Field) {
	res = &Field{
		Type: field.Type,
	}
	if len(field.Names) < 1 {
		return
	}
	if field.Names[0] == nil {
		return
	}
	res.Name = field.Names[0].Name
	return
}
