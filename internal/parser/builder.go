package parser

import (
	"github.com/skhoroshavin/automap/internal/utils"
	"go/ast"
)

func findBuilderExpr(fun *ast.FuncDecl) *ast.CallExpr {
	// Mapper body is always one statement, which should be an expression
	if len(fun.Body.List) != 1 {
		return nil
	}
	expr, ok := fun.Body.List[0].(*ast.ExprStmt)
	if !ok {
		return nil
	}

	// This expression should be a panic call with a single argument
	panicCall, ok := expr.X.(*ast.CallExpr)
	if !ok {
		return nil
	}
	if utils.AST2String(panicCall.Fun) != "panic" {
		return nil
	}
	if len(panicCall.Args) != 1 {
		return nil
	}

	// This argument should be call to automap.Build
	buildCall, ok := panicCall.Args[0].(*ast.CallExpr)
	if !ok {
		return nil
	}
	if utils.AST2String(buildCall.Fun) != "automap.Build" {
		return nil
	}

	return buildCall
}
