package parser

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

func nodeToString(expr ast.Node) string {
	buf := bytes.Buffer{}
	fset := token.NewFileSet()
	err := printer.Fprint(&buf, fset, expr)
	if err != nil {
		return ""
	}
	return buf.String()
}
