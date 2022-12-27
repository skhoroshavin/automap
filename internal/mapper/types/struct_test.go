package types

import (
	"github.com/skhoroshavin/automap/internal/mapper/node"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestStruct(t *testing.T) {
	suite.Run(t, new(StructSuite))
}

type StructSuite struct {
	suite.Suite
}

var StringType = &Opaque{Name_: "string"}

func (s *StructSuite) TestDirectMapping() {
	strct := &Struct{
		Name_: "Answer",
		Fields: ProviderList{
			{Name: "Value", Type: StringType},
			{Name: "Question", Type: StringType},
		},
		Getters: ProviderList{},
	}
	mapper, err := strct.BuildMapper([]Provider{
		{Name: "answer", Type: strct},
	})

	s.Assert().NoError(err)
	s.Assert().Equal(node.NewValue("answer"), mapper)
}

func (s *StructSuite) TestArgsMapping() {
	strct := &Struct{
		Name_: "Answer",
		Fields: ProviderList{
			{Name: "Value", Type: StringType},
			{Name: "Question", Type: StringType},
		},
		Getters: ProviderList{},
	}
	mapper, err := strct.BuildMapper(ProviderList{
		{Name: "question", Type: StringType},
		{Name: "value", Type: StringType},
	})

	s.Assert().NoError(err)
	s.Assert().Equal(node.NewStruct(
		"Answer",
		node.NewField("Value", node.NewValue("value")),
		node.NewField("Question", node.NewValue("question")),
	), mapper)
}

func (s *StructSuite) TestSimpleStructMapping() {
	target := &Struct{
		Name_: "Answer",
		Fields: ProviderList{
			{Name: "Value", Type: StringType},
			{Name: "Question", Type: StringType},
		},
		Getters: ProviderList{},
	}
	source := &Struct{
		Name_: "mapper.Answer",
		Fields: ProviderList{
			{Name: "Value", Type: StringType},
			{Name: "Question", Type: StringType},
			{Name: "Reason", Type: StringType},
		},
		Getters: ProviderList{},
	}
	mapper, err := target.BuildMapper(ProviderList{
		{Name: "v", Type: source},
	})

	s.Assert().NoError(err)
	s.Assert().Equal(node.NewStruct(
		"Answer",
		node.NewField("Value", node.NewValue("v.Value")),
		node.NewField("Question", node.NewValue("v.Question")),
	), mapper)
}
