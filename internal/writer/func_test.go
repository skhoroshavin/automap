package writer

import (
	"bytes"
	"github.com/skhoroshavin/automap/internal/core"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestWriteFunc(t *testing.T) {
	suite.Run(t, new(WriteFuncSuite))
}

type WriteFuncSuite struct {
	suite.Suite
	out *bytes.Buffer
}

func (s *WriteFuncSuite) SetupTest() {
	s.out = &bytes.Buffer{}
}

func (s *WriteFuncSuite) TestWriteEmptyFuncFails() {
	err := writeFunc(s.out, &core.FuncBody{})
	s.Assert().Error(err)
}

func (s *WriteFuncSuite) TestWriteFuncWithSingleResultCreatesSingleReturnStatement() {
	err := writeFunc(s.out, &core.FuncBody{
		Result: "42",
	})
	s.Assert().NoError(err)
	s.Assert().Equal("\treturn 42\n", s.out.String())
}

func (s *WriteFuncSuite) TestWriteFuncWithoutResultFails() {
	err := writeFunc(s.out, &core.FuncBody{
		Vars: []core.Variable{
			{Name: "tmp", Value: "42"},
		},
	})
	s.Assert().Error(err)
}

func (s *WriteFuncSuite) TestWriteFuncWithVariablesCreatesThemInOrderThenReturnStatement() {
	err := writeFunc(s.out, &core.FuncBody{
		Vars: []core.Variable{
			{Name: "a", Value: "20"},
			{Name: "b", Value: "a + 2"},
		},
		Result: "a + b",
	})
	s.Assert().NoError(err)

	expected :=
		`	a := 20
	b := a + 2
	return a + b
`
	s.Assert().Equal(expected, s.out.String())
}
