package provider

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestProviderList(t *testing.T) {
	suite.Run(t, new(ProviderListSuite))
}

type ProviderListSuite struct {
	suite.Suite
	deepList     List
	visitedNodes string
}

func (s *ProviderListSuite) SetupSuite() {
	s.deepList = List{
		NewMock("A",
			NewMock("C",
				NewMock("G"),
				NewMock("H"),
			),
			NewMock("D",
				NewMock("I"),
				NewMock("J"),
			),
		),
		NewMock("B",
			NewMock("E",
				NewMock("K"),
				NewMock("L"),
			),
			NewMock("F",
				NewMock("M"),
				NewMock("N"),
			),
		),
	}
}

func (s *ProviderListSuite) SetupTest() {
	s.visitedNodes = ""
}

func (s *ProviderListSuite) TestForEachEmpty() {
	List{}.ForEach(s.visitAll())

	s.Assert().Empty(s.visitedNodes)
}

func (s *ProviderListSuite) TestForEachFlat() {
	List{
		NewMock("A"),
		NewMock("B"),
		NewMock("C"),
	}.ForEach(s.visitAll())

	s.Assert().Equal("ABC", s.visitedNodes)
}

func (s *ProviderListSuite) TestForEachNested() {
	List{
		NewMock("A",
			NewMock("D"),
			NewMock("E"),
		),
		NewMock("B",
			NewMock("F"),
			NewMock("G"),
		),
		NewMock("C",
			NewMock("H"),
			NewMock("I"),
		),
	}.ForEach(s.visitAll())

	s.Assert().Equal("ABCDEFGHI", s.visitedNodes)
}

func (s *ProviderListSuite) TestForEachDeepNested() {
	s.deepList.ForEach(s.visitAll())

	s.Assert().Equal("ABCDEFGHIJKLMN", s.visitedNodes)
}

func (s *ProviderListSuite) TestForEachStopShallow() {
	s.deepList.ForEach(s.visitUntil("B"))

	s.Assert().Equal("AB", s.visitedNodes)
}

func (s *ProviderListSuite) TestForEachStopDeep() {
	s.deepList.ForEach(s.visitUntil("I"))

	s.Assert().Equal("ABCDEFGHI", s.visitedNodes)
}

func (s *ProviderListSuite) visitAll() Visitor {
	return func(p Provider) bool {
		s.visitedNodes = s.visitedNodes + p.(*Mock).name
		return false
	}
}

func (s *ProviderListSuite) visitUntil(name string) Visitor {
	return func(p Provider) bool {
		pName := p.(*Mock).name
		s.visitedNodes = s.visitedNodes + pName
		return pName == name
	}
}
