package writer

import (
	"bytes"
	"github.com/skhoroshavin/automap/internal/core/ast"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestWriteExpr(t *testing.T) {
	suite.Run(t, new(WriteExprSuite))
}

type WriteExprSuite struct {
	suite.Suite
	out *bytes.Buffer
}

func (s *WriteExprSuite) SetupTest() {
	s.out = &bytes.Buffer{}
}

func (s *WriteExprSuite) TestWriteValueExpr() {
	err := writeExpr(s.out, ast.NewValue("42"), 1)
	s.Assert().NoError(err)

	s.Assert().Equal("42", s.out.String())
}

func (s *WriteExprSuite) TestStructExpr() {
	err := writeExpr(s.out, ast.NewStruct(
		"Answer",
		ast.NewField("Question", ast.NewValue("\"wtf\"")),
		ast.NewField("Value", ast.NewValue("42")),
	), 1)
	s.Assert().NoError(err)

	expected := `Answer{
		Question: "wtf",
		Value: 42,
	}`
	s.Assert().Equal(expected, s.out.String())
}

func (s *WriteExprSuite) TestPointerStructExpr() {
	err := writeExpr(s.out, ast.NewStructPtr(
		"Answer",
		ast.NewField("Question", ast.NewValue("\"wtf\"")),
		ast.NewField("Value", ast.NewValue("42")),
	), 1)
	s.Assert().NoError(err)

	expected := `&Answer{
		Question: "wtf",
		Value: 42,
	}`
	s.Assert().Equal(expected, s.out.String())
}

func (s *WriteExprSuite) TestNestedStructExpr() {
	err := writeExpr(s.out, &ast.StructExpr{
		Name: "Question",
		Fields: []ast.Field{
			ast.NewField("Value", ast.NewValue("\"wtf\"")),
			ast.NewField("Answer", ast.NewStructPtr(
				"Answer",
				ast.NewField("Value", ast.NewValue("42")),
			)),
		},
	}, 1)
	s.Assert().NoError(err)

	expected := `Question{
		Value: "wtf",
		Answer: &Answer{
			Value: 42,
		},
	}`
	s.Assert().Equal(expected, s.out.String())
}
