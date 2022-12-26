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

var StringType = &OpaqueType{Name_: "string"}

func (s *StructSuite) TestDirectMapping() {
	strct := &StructType{
		Name_: "Answer",
		Fields: ProviderList{
			{Name: "Value", Type: StringType},
			{Name: "Question", Type: StringType},
		},
		Getters: ProviderList{},
	}
	mapper := strct.BuildMapper([]Provider{
		{Name: "answer", Type: strct},
	})

	s.Assert().Equal(&ValueNode{Value: "answer"}, mapper)
}

func (s *StructSuite) TestArgsMapping() {
	strct := &StructType{
		Name_: "Answer",
		Fields: ProviderList{
			{Name: "Value", Type: StringType},
			{Name: "Question", Type: StringType},
		},
		Getters: ProviderList{},
	}
	mapper := strct.BuildMapper(ProviderList{
		{Name: "question", Type: StringType},
		{Name: "value", Type: StringType},
	})

	s.Assert().Equal(&StructNode{
		Name: "Answer",
		Fields: []NamedNode{
			{Name: "Value", Value: &ValueNode{Value: "value"}},
			{Name: "Question", Value: &ValueNode{Value: "question"}},
		},
	}, mapper)
}

func (s *StructSuite) TestSimpleStructMapping() {
	target := &StructType{
		Name_: "Answer",
		Fields: ProviderList{
			{Name: "Value", Type: StringType},
			{Name: "Question", Type: StringType},
		},
		Getters: ProviderList{},
	}
	source := &StructType{
		Name_: "core.Answer",
		Fields: ProviderList{
			{Name: "Value", Type: StringType},
			{Name: "Question", Type: StringType},
			{Name: "Reason", Type: StringType},
		},
		Getters: ProviderList{},
	}
	mapper := target.BuildMapper(ProviderList{
		{Name: "v", Type: source},
	})

	s.Assert().Equal(&StructNode{
		Name: "Answer",
		Fields: []NamedNode{
			{Name: "Value", Value: &ValueNode{Value: "v.Value"}},
			{Name: "Question", Value: &ValueNode{Value: "v.Question"}},
		},
	}, mapper)
}
