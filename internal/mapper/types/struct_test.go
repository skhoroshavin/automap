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

func (s *StructSuite) TestDirectMapping() {
	strct := &Struct{
		name: "Answer",
		fields: ProviderList{
			{Name: "Value", Type: NewOpaque("string")},
			{Name: "Question", Type: NewOpaque("string")},
		},
		getters: ProviderList{},
	}
	mapper, err := strct.BuildMapper([]Provider{
		{Name: "answer", Type: strct},
	})

	s.Assert().NoError(err)
	s.Assert().Equal(node.NewValue("answer"), mapper)
}

func (s *StructSuite) TestArgsMapping() {
	strct := &Struct{
		name: "Answer",
		fields: ProviderList{
			{Name: "Value", Type: NewOpaque("string")},
			{Name: "Question", Type: NewOpaque("string")},
		},
		getters: ProviderList{},
	}
	mapper, err := strct.BuildMapper(ProviderList{
		{Name: "question", Type: NewOpaque("string")},
		{Name: "value", Type: NewOpaque("string")},
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
		name: "Answer",
		fields: ProviderList{
			{Name: "Value", Type: NewOpaque("string")},
			{Name: "Question", Type: NewOpaque("string")},
		},
		getters: ProviderList{},
	}
	source := &Struct{
		name: "mapper.Answer",
		fields: ProviderList{
			{Name: "Value", Type: NewOpaque("string")},
			{Name: "Question", Type: NewOpaque("string")},
			{Name: "Reason", Type: NewOpaque("string")},
		},
		getters: ProviderList{},
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
