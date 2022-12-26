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
	mapper *ast.Mapper
}

func (s *StructSuite) SetupTest() {
	s.mapper = new(ast.Mapper)
}

func (s *StructSuite) TestStructBuildsSimpleReturnStatement() {
	NewStruct(
		"Answer",
		NewField("Question", NewValue("\"wtf\"")),
		NewField("Value", NewValue("42")),
	).CompileTo(s.mapper)

	s.Assert().Empty(s.mapper.Vars)

	expected := ast.NewStruct(
		"Answer",
		ast.NewField("Question", ast.NewValue("\"wtf\"")),
		ast.NewField("Value", ast.NewValue("42")),
	)
	s.Assert().Equal(expected, s.mapper.Result)
}

func (s *StructSuite) TestPointerStructBuildsSimpleReturnStatement() {
	NewStructPtr(
		"Answer",
		NewField("Question", NewValue("\"wtf\"")),
		NewField("Value", NewValue("42")),
	).CompileTo(s.mapper)

	s.Assert().Empty(s.mapper.Vars)

	expected := ast.NewStructPtr(
		"Answer",
		ast.NewField("Question", ast.NewValue("\"wtf\"")),
		ast.NewField("Value", ast.NewValue("42")),
	)
	s.Assert().Equal(expected, s.mapper.Result)
}

func (s *StructSuite) TestNestedStructBuildsSimpleReturnStatement() {
	NewStruct(
		"Question",
		NewField("Value", NewValue("\"wtf\"")),
		NewField("Answer", NewStructPtr(
			"Answer",
			NewField("Value", NewValue("42")),
		)),
	).CompileTo(s.mapper)

	s.Assert().Empty(s.mapper.Vars)

	expected := ast.NewStruct(
		"Question",
		ast.NewField("Value", ast.NewValue("\"wtf\"")),
		ast.NewField("Answer", ast.NewStructPtr(
			"Answer",
			ast.NewField("Value", ast.NewValue("42")),
		)),
	)
	s.Assert().Equal(expected, s.mapper.Result)
}
