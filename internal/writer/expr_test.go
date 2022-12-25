package writer

import (
	"bytes"
	"github.com/skhoroshavin/automap/internal/core"
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
	err := writeExpr(s.out, &core.ValueExpr{Value: "42"}, 1)
	s.Assert().NoError(err)

	s.Assert().Equal("42", s.out.String())
}

func (s *WriteExprSuite) TestStructExpr() {
	err := writeExpr(s.out, &core.StructExpr{
		Name: "Answer",
		Fields: []core.FieldExpr{
			{Name: "Question", Value: &core.ValueExpr{Value: "\"wtf\""}},
			{Name: "Value", Value: &core.ValueExpr{Value: "42"}},
		},
	}, 1)
	s.Assert().NoError(err)

	expected := `Answer{
		Question: "wtf",
		Value: 42,
	}`
	s.Assert().Equal(expected, s.out.String())
}

func (s *WriteExprSuite) TestPointerStructExpr() {
	err := writeExpr(s.out, &core.StructExpr{
		Name:      "Answer",
		IsPointer: true,
		Fields: []core.FieldExpr{
			{Name: "Question", Value: &core.ValueExpr{Value: "\"wtf\""}},
			{Name: "Value", Value: &core.ValueExpr{Value: "42"}},
		},
	}, 1)
	s.Assert().NoError(err)

	expected := `&Answer{
		Question: "wtf",
		Value: 42,
	}`
	s.Assert().Equal(expected, s.out.String())
}

func (s *WriteExprSuite) TestNestedStructExpr() {
	err := writeExpr(s.out, &core.StructExpr{
		Name: "Question",
		Fields: []core.FieldExpr{
			{Name: "Value", Value: &core.ValueExpr{Value: "\"wtf\""}},
			{Name: "Answer", Value: &core.StructExpr{
				Name:      "Answer",
				IsPointer: true,
				Fields: []core.FieldExpr{
					{Name: "Value", Value: &core.ValueExpr{Value: "42"}},
				},
			}},
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
