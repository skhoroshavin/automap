package parser

import (
	"github.com/stretchr/testify/suite"
	"go/ast"
	"testing"
)

func TestImports(t *testing.T) {
	suite.Run(t, new(ImportsSuite))
}

type ImportsSuite struct {
	suite.Suite
}

func (s *ImportsSuite) TestTrivial() {
	path, name := parseImport(&ast.ImportSpec{
		Path: &ast.BasicLit{Value: `"strings"`},
	})

	s.Assert().Equal("strings", path)
	s.Assert().Equal("strings", name)
}

func (s *ImportsSuite) TestUnnamed() {
	path, name := parseImport(&ast.ImportSpec{
		Path: &ast.BasicLit{Value: `"github.com/skhoroshavin/automap"`},
	})

	s.Assert().Equal("github.com/skhoroshavin/automap", path)
	s.Assert().Equal("automap", name)
}

func (s *ImportsSuite) TestNamed() {
	path, name := parseImport(&ast.ImportSpec{
		Name: &ast.Ident{Name: "xxx"},
		Path: &ast.BasicLit{Value: `"github.com/skhoroshavin/automap"`},
	})

	s.Assert().Equal("github.com/skhoroshavin/automap", path)
	s.Assert().Equal("xxx", name)
}

func (s *ImportsSuite) TestConvertToList() {
	imports := ImportMap{
		"string":                          "string",
		"github.com/skhoroshavin/automap": "automap",
		"github.com/skhoroshavin/automap/internal": "impl",
	}

	s.Assert().ElementsMatch([]string{
		`"string"`,
		`"github.com/skhoroshavin/automap"`,
		`impl "github.com/skhoroshavin/automap/internal"`,
	}, imports.ToList())
}
