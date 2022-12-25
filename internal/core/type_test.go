package core

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestStruct(t *testing.T) {
	suite.Run(t, new(StructSuite))
}

type StructSuite struct {
	suite.Suite
}

func (s *StructSuite) TestDirectMapping() {
	strct := &StructType{
		Name: "Answer",
		Fields: []Provider{
			{Name: "Value", Type: "string"},
			{Name: "Question", Type: "string"},
		},
		Getters: []Provider{},
	}
	mapper := strct.BuildMapper([]Provider{
		{Name: "answer", Type: "Answer"},
	})

	s.Assert().Equal(&ValueNode{Value: "answer"}, mapper)
}

func (s *StructSuite) TestArgsMapping() {
	strct := &StructType{
		Name: "Answer",
		Fields: []Provider{
			{Name: "Value", Type: "string"},
			{Name: "Question", Type: "string"},
		},
		Getters: []Provider{},
	}
	mapper := strct.BuildMapper([]Provider{
		{Name: "question", Type: "string"},
		{Name: "value", Type: "string"},
	})

	s.Assert().Equal(&StructNode{
		Name: "Answer",
		Fields: []NamedNode{
			{Name: "Value", Value: &ValueNode{Value: "value"}},
			{Name: "Question", Value: &ValueNode{Value: "question"}},
		},
	}, mapper)
}
