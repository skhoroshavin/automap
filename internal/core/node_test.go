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

func (s *NodeSuite) TestExprBuildsSimpleReturnStatement() {
	body := BuildFuncBody(&ExprNode{Value: "42"})

	s.Assert().Empty(body.Vars)
	s.Assert().Equal("42", body.Result)
}

func (s *NodeSuite) TestStructBuildsSimpleReturnStatement() {
	body := BuildFuncBody(&StructNode{
		Name: "Answer",
		Fields: []NamedNode{
			{Name: "Question", Value: &ExprNode{Value: "\"wtf\""}},
			{Name: "Value", Value: &ExprNode{Value: "42"}},
		},
	})

	s.Assert().Empty(body.Vars)

	expected := `Answer{
		Question: "wtf",
		Value: 42,
	}`
	s.Assert().Equal(expected, body.Result)
}

func (s *NodeSuite) TestPointerStructBuildsSimpleReturnStatement() {
	body := BuildFuncBody(&StructNode{
		Name:      "Answer",
		IsPointer: true,
		Fields: []NamedNode{
			{Name: "Question", Value: &ExprNode{Value: "\"wtf\""}},
			{Name: "Value", Value: &ExprNode{Value: "42"}},
		},
	})

	s.Assert().Empty(body.Vars)

	expected := `&Answer{
		Question: "wtf",
		Value: 42,
	}`
	s.Assert().Equal(expected, body.Result)
}

func (s *NodeSuite) TestNestedStructBuildsSimpleReturnStatement() {
	body := BuildFuncBody(&StructNode{
		Name: "Question",
		Fields: []NamedNode{
			{Name: "Value", Value: &ExprNode{Value: "\"wtf\""}},
			{Name: "Answer", Value: &StructNode{
				Name:      "Answer",
				IsPointer: true,
				Fields: []NamedNode{
					{Name: "Value", Value: &ExprNode{Value: "42"}},
				},
			}},
		},
	})

	s.Assert().Empty(body.Vars)

	expected := `Question{
		Value: "wtf",
		Answer: &Answer{
			Value: 42,
		},
	}`
	s.Assert().Equal(expected, body.Result)
}
