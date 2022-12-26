package writer

import (
	"bytes"
	"github.com/skhoroshavin/automap/internal/core/ast"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestWriteFunc(t *testing.T) {
	suite.Run(t, new(WriteFuncSuite))
}

type WriteFuncSuite struct {
	suite.Suite
	out *bytes.Buffer
}

func (s *WriteFuncSuite) SetupTest() {
	s.out = &bytes.Buffer{}
}

func (s *WriteFuncSuite) TestWriteEmptyFuncFails() {
	err := writeFunc(s.out, &ast.Mapper{})
	s.Assert().Error(err)
}

func (s *WriteFuncSuite) TestWriteFuncWithSingleResultCreatesSingleReturnStatement() {
	err := writeFunc(s.out, &ast.Mapper{
		Result: &ast.ValueExpr{Value: "42"},
	})
	s.Assert().NoError(err)
	s.Assert().Equal("\treturn 42\n", s.out.String())
}

func (s *WriteFuncSuite) TestWriteFuncWithoutResultFails() {
	err := writeFunc(s.out, &ast.Mapper{
		Vars: []ast.Variable{
			{Name: "tmp", Value: &ast.ValueExpr{Value: "42"}},
		},
	})
	s.Assert().Error(err)
}

func (s *WriteFuncSuite) TestWriteFuncWithVariablesCreatesThemInOrderThenReturnStatement() {
	err := writeFunc(s.out, &ast.Mapper{
		Vars: []ast.Variable{
			{Name: "a", Value: &ast.ValueExpr{Value: "20"}},
			{Name: "b", Value: &ast.ValueExpr{Value: "a + 2"}},
		},
		Result: &ast.ValueExpr{Value: "a + b"},
	})
	s.Assert().NoError(err)

	expected :=
		`	a := 20
	b := a + 2
	return a + b
`
	s.Assert().Equal(expected, s.out.String())
}

func (s *WriteFuncSuite) TestWriteFuncHandlesComplexExpressions() {
	err := writeFunc(s.out, &ast.Mapper{
		Vars: []ast.Variable{
			{Name: "answer", Value: &ast.StructExpr{
				Name:      "Answer",
				IsPointer: true,
				Fields: []ast.Field{
					{Name: "Value", Value: &ast.ValueExpr{Value: "42"}},
				},
			}},
		},
		Result: &ast.StructExpr{
			Name: "Question",
			Fields: []ast.Field{
				{Name: "Value", Value: &ast.ValueExpr{Value: "\"wtf\""}},
				{Name: "Answer", Value: &ast.ValueExpr{Value: "answer"}},
			},
		},
	})
	s.Assert().NoError(err)

	expected := `	answer := &Answer{
		Value: 42,
	}
	return Question{
		Value: "wtf",
		Answer: answer,
	}
`
	s.Assert().Equal(expected, s.out.String())
}
