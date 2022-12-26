package writer

import (
	"bytes"
	"github.com/skhoroshavin/automap/internal/core/ast"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestWriteMapper(t *testing.T) {
	suite.Run(t, new(WriteMapperSuite))
}

type WriteMapperSuite struct {
	suite.Suite
	out *bytes.Buffer
}

func (s *WriteMapperSuite) SetupTest() {
	s.out = &bytes.Buffer{}
}

func (s *WriteMapperSuite) TestWriteEmptyMapperFails() {
	err := writeMapper(s.out, &ast.Mapper{})
	s.Assert().Error(err)
}

func (s *WriteMapperSuite) TestWriteMapperWithSimpleReturnStatement() {
	err := writeMapper(s.out, &ast.Mapper{
		Signature: "func GetAnswer() int",
		Result:    ast.NewValue("42"),
	})
	s.Assert().NoError(err)

	expected := `func GetAnswer() int {
	return 42
}
`
	s.Assert().Equal(expected, s.out.String())
}

func (s *WriteMapperSuite) TestWriteMapperWithoutReturnStatementFails() {
	err := writeMapper(s.out, &ast.Mapper{
		Signature: "func GetAnswer() int",
		Vars: []ast.Variable{
			{Name: "tmp", Value: ast.NewValue("42")},
		},
	})
	s.Assert().Error(err)
}

func (s *WriteMapperSuite) TestWriteMapperWithVariablesCreatesThemInOrderThenReturnStatement() {
	err := writeMapper(s.out, &ast.Mapper{
		Signature: "func GetAnswer() int",
		Vars: []ast.Variable{
			{Name: "a", Value: ast.NewValue("20")},
			{Name: "b", Value: ast.NewValue("a + 2")},
		},
		Result: ast.NewValue("a + b"),
	})
	s.Assert().NoError(err)

	expected := `func GetAnswer() int {
	a := 20
	b := a + 2
	return a + b
}
`
	s.Assert().Equal(expected, s.out.String())
}

func (s *WriteMapperSuite) TestWriteMapperHandlesComplexExpressions() {
	err := writeMapper(s.out, &ast.Mapper{
		Signature: "func GetQuestion() Question",
		Vars: []ast.Variable{{
			Name: "answer",
			Value: ast.NewStructPtr(
				"Answer",
				ast.NewField("Value", ast.NewValue("42")),
			)},
		},
		Result: ast.NewStruct(
			"Question",
			ast.NewField("Value", ast.NewValue("\"wtf\"")),
			ast.NewField("Answer", ast.NewValue("answer")),
		),
	})
	s.Assert().NoError(err)

	expected := `func GetQuestion() Question {
	answer := &Answer{
		Value: 42,
	}
	return Question{
		Value: "wtf",
		Answer: answer,
	}
}
`
	s.Assert().Equal(expected, s.out.String())
}
