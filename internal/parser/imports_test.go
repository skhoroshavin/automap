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
	imports Imports
}

func (s *ImportsSuite) SetupTest() {
	s.imports = NewImports()
}

func (s *ImportsSuite) TestTrivial() {
	s.imports.ParseImport(&ast.ImportSpec{
		Path: &ast.BasicLit{Value: `"strings"`},
	})

	s.Assert().Equal(Imports{"strings": "strings"}, s.imports)
}

func (s *ImportsSuite) TestUnnamed() {
	s.imports.ParseImport(&ast.ImportSpec{
		Path: &ast.BasicLit{Value: `"github.com/skhoroshavin/automap/internal/mapper"`},
	})

	s.Assert().Equal(Imports{"github.com/skhoroshavin/automap/internal/mapper": "mapper"}, s.imports)
}

func (s *ImportsSuite) TestNamed() {
	s.imports.ParseImport(&ast.ImportSpec{
		Name: &ast.Ident{Name: "impl"},
		Path: &ast.BasicLit{Value: `"github.com/skhoroshavin/automap/internal"`},
	})

	s.Assert().Equal(Imports{"github.com/skhoroshavin/automap/internal": "impl"}, s.imports)
}

func (s *ImportsSuite) TestSkipAutomap() {
	s.imports.ParseImport(&ast.ImportSpec{
		Path: &ast.BasicLit{Value: `"github.com/skhoroshavin/automap"`},
	})

	s.Assert().Empty(s.imports)
}

func (s *ImportsSuite) TestConvertToList() {
	imports := Imports{
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
