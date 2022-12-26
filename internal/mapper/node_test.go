package mapper

import (
	"github.com/skhoroshavin/automap/internal/mapper/ast"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestNode(t *testing.T) {
	suite.Run(t, new(NodeSuite))
}

type NodeSuite struct {
	suite.Suite
}

func (s *NodeSuite) TestValueBuildsSimpleReturnStatement() {
	body := BuildFuncBody(&ValueNode{Value: "42"})

	s.Assert().Empty(body.Vars)
	s.Assert().Equal(ast.NewValue("42"), body.Result)
}

func (s *NodeSuite) TestStructBuildsSimpleReturnStatement() {
	body := BuildFuncBody(&StructNode{
		Name: "Answer",
		Fields: []NamedNode{
			{Name: "Question", Value: &ValueNode{Value: "\"wtf\""}},
			{Name: "Value", Value: &ValueNode{Value: "42"}},
		},
	})

	s.Assert().Empty(body.Vars)

	expected := ast.NewStruct(
		"Answer",
		ast.NewField("Question", ast.NewValue("\"wtf\"")),
		ast.NewField("Value", ast.NewValue("42")),
	)
	s.Assert().Equal(expected, body.Result)
}

func (s *NodeSuite) TestPointerStructBuildsSimpleReturnStatement() {
	body := BuildFuncBody(&StructNode{
		Name:      "Answer",
		IsPointer: true,
		Fields: []NamedNode{
			{Name: "Question", Value: &ValueNode{Value: "\"wtf\""}},
			{Name: "Value", Value: &ValueNode{Value: "42"}},
		},
	})

	s.Assert().Empty(body.Vars)

	expected := ast.NewStructPtr(
		"Answer",
		ast.NewField("Question", ast.NewValue("\"wtf\"")),
		ast.NewField("Value", ast.NewValue("42")),
	)
	s.Assert().Equal(expected, body.Result)
}

func (s *NodeSuite) TestNestedStructBuildsSimpleReturnStatement() {
	body := BuildFuncBody(&StructNode{
		Name: "Question",
		Fields: []NamedNode{
			{Name: "Value", Value: &ValueNode{Value: "\"wtf\""}},
			{Name: "Answer", Value: &StructNode{
				Name:      "Answer",
				IsPointer: true,
				Fields: []NamedNode{
					{Name: "Value", Value: &ValueNode{Value: "42"}},
				},
			}},
		},
	})

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
