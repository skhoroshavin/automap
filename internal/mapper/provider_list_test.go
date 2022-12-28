package mapper

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestProviderList(t *testing.T) {
	suite.Run(t, new(ProviderListSuite))
}

type ProviderListSuite struct {
	suite.Suite
	visitedNodes string
}

func (s *ProviderListSuite) SetupTest() {
	s.visitedNodes = ""
}

func (s *ProviderListSuite) TestForEachEmpty() {
	ProviderList{}.ForEach(s.visitAll())

	s.Assert().Empty(s.visitedNodes)
}

func (s *ProviderListSuite) TestForEachFlat() {
	ProviderList{
		&MockProvider{name: "A"},
		&MockProvider{name: "B"},
		&MockProvider{name: "C"},
	}.ForEach(s.visitAll())

	s.Assert().Equal("ABC", s.visitedNodes)
}

func (s *ProviderListSuite) TestForEachNested() {
	ProviderList{
		&MockProvider{name: "A", children: ProviderList{
			&MockProvider{name: "D"},
			&MockProvider{name: "E"},
		}},
		&MockProvider{name: "B", children: ProviderList{
			&MockProvider{name: "F"},
			&MockProvider{name: "G"},
		}},
		&MockProvider{name: "C", children: ProviderList{
			&MockProvider{name: "H"},
			&MockProvider{name: "I"},
		}},
	}.ForEach(s.visitAll())

	s.Assert().Equal("ABCDEFGHI", s.visitedNodes)
}

func (s *ProviderListSuite) TestForEachDeepNested() {
	ProviderList{
		&MockProvider{name: "A", children: ProviderList{
			&MockProvider{name: "C", children: ProviderList{
				&MockProvider{name: "G"},
				&MockProvider{name: "H"},
			}},
			&MockProvider{name: "D", children: ProviderList{
				&MockProvider{name: "I"},
				&MockProvider{name: "J"},
			}},
		}},
		&MockProvider{name: "B", children: ProviderList{
			&MockProvider{name: "E", children: ProviderList{
				&MockProvider{name: "K"},
				&MockProvider{name: "L"},
			}},
			&MockProvider{name: "F", children: ProviderList{
				&MockProvider{name: "M"},
				&MockProvider{name: "N"},
			}},
		}},
	}.ForEach(s.visitAll())

	s.Assert().Equal("ABCDEFGHIJKLMN", s.visitedNodes)
}

func (s *ProviderListSuite) TestForEachStopShallow() {
	ProviderList{
		&MockProvider{name: "A", children: ProviderList{
			&MockProvider{name: "C", children: ProviderList{
				&MockProvider{name: "G"},
				&MockProvider{name: "H"},
			}},
			&MockProvider{name: "D", children: ProviderList{
				&MockProvider{name: "I"},
				&MockProvider{name: "J"},
			}},
		}},
		&MockProvider{name: "B", children: ProviderList{
			&MockProvider{name: "E", children: ProviderList{
				&MockProvider{name: "K"},
				&MockProvider{name: "L"},
			}},
			&MockProvider{name: "F", children: ProviderList{
				&MockProvider{name: "M"},
				&MockProvider{name: "N"},
			}},
		}},
	}.ForEach(s.visitUntil("B"))

	s.Assert().Equal("AB", s.visitedNodes)
}

func (s *ProviderListSuite) TestForEachStopDeep() {
	ProviderList{
		&MockProvider{name: "A", children: ProviderList{
			&MockProvider{name: "C", children: ProviderList{
				&MockProvider{name: "G"},
				&MockProvider{name: "H"},
			}},
			&MockProvider{name: "D", children: ProviderList{
				&MockProvider{name: "I"},
				&MockProvider{name: "J"},
			}},
		}},
		&MockProvider{name: "B", children: ProviderList{
			&MockProvider{name: "E", children: ProviderList{
				&MockProvider{name: "K"},
				&MockProvider{name: "L"},
			}},
			&MockProvider{name: "F", children: ProviderList{
				&MockProvider{name: "M"},
				&MockProvider{name: "N"},
			}},
		}},
	}.ForEach(s.visitUntil("I"))

	s.Assert().Equal("ABCDEFGHI", s.visitedNodes)
}

func (s *ProviderListSuite) visitAll() ProviderVisitor {
	return func(p Provider) bool {
		s.visitedNodes = s.visitedNodes + p.Name()
		return false
	}
}

func (s *ProviderListSuite) visitUntil(name string) ProviderVisitor {
	return func(p Provider) bool {
		s.visitedNodes = s.visitedNodes + p.Name()
		return p.Name() == name
	}
}
