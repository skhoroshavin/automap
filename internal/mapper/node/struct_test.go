package node

import (
	"github.com/skhoroshavin/automap/internal/mapper/ast"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestStruct(t *testing.T) {
	suite.Run(t, new(StructSuite))
}

type StructSuite struct {
	suite.Suite
}

func (s *StructSuite) TestStructBuildsSimpleReturnStatement() {
	body := build(NewStruct(
		"Answer",
		NewField("Question", NewValue("\"wtf\"")),
		NewField("Value", NewValue("42")),
	))

	s.Assert().Empty(body.Vars)

	expected := ast.NewStruct(
		"Answer",
		ast.NewField("Question", ast.NewValue("\"wtf\"")),
		ast.NewField("Value", ast.NewValue("42")),
	)
	s.Assert().Equal(expected, body.Result)
}

func (s *StructSuite) TestPointerStructBuildsSimpleReturnStatement() {
	body := build(NewStructPtr(
		"Answer",
		NewField("Question", NewValue("\"wtf\"")),
		NewField("Value", NewValue("42")),
	))

	s.Assert().Empty(body.Vars)

	expected := ast.NewStructPtr(
		"Answer",
		ast.NewField("Question", ast.NewValue("\"wtf\"")),
		ast.NewField("Value", ast.NewValue("42")),
	)
	s.Assert().Equal(expected, body.Result)
}

func (s *StructSuite) TestNestedStructBuildsSimpleReturnStatement() {
	body := build(NewStruct(
		"Question",
		NewField("Value", NewValue("\"wtf\"")),
		NewField("Answer", NewStructPtr(
			"Answer",
			NewField("Value", NewValue("42")),
		)),
	))

	s.Assert().Empty(body.Vars)

	expected := ast.NewStruct(
		"Question",
		ast.NewField("Value", ast.NewValue("\"wtf\"")),
		ast.NewField("Answer", ast.NewStructPtr(
			"Answer",
			ast.NewField("Value", ast.NewValue("42")),
		)),
	)
	s.Assert().Equal(expected, body.Result)
}
