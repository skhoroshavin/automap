package utils

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

func AST2String(expr ast.Node) string {
	buf := bytes.Buffer{}
	fset := token.NewFileSet()
	err := printer.Fprint(&buf, fset, expr)
	if err != nil {
		return ""
	}
	return buf.String()
}
