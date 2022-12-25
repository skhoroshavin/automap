package core

import (
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
	s.Assert().Equal(&ValueExpr{Value: "42"}, body.Result)
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

	expected := &StructExpr{
		Name: "Answer",
		Fields: []FieldExpr{
			{Name: "Question", Value: &ValueExpr{Value: "\"wtf\""}},
			{Name: "Value", Value: &ValueExpr{Value: "42"}},
		},
	}
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

	expected := &StructExpr{
		Name:      "Answer",
		IsPointer: true,
		Fields: []FieldExpr{
			{Name: "Question", Value: &ValueExpr{Value: "\"wtf\""}},
			{Name: "Value", Value: &ValueExpr{Value: "42"}},
		},
	}
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

	expected := &StructExpr{
		Name: "Question",
		Fields: []FieldExpr{
			{Name: "Value", Value: &ValueExpr{Value: "\"wtf\""}},
			{Name: "Answer", Value: &StructExpr{
				Name:      "Answer",
				IsPointer: true,
				Fields: []FieldExpr{
					{Name: "Value", Value: &ValueExpr{Value: "42"}},
				},
			}},
		},
	}
	s.Assert().Equal(expected, body.Result)
}
