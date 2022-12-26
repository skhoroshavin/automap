package node

import (
	"github.com/skhoroshavin/automap/internal/mapper/ast"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestValue(t *testing.T) {
	suite.Run(t, new(ValueSuite))
}

type ValueSuite struct {
	suite.Suite
	mapper *ast.Mapper
}

func (s *ValueSuite) SetupTest() {
	s.mapper = new(ast.Mapper)
}

func (s *ValueSuite) TestValueBuildsSimpleReturnStatement() {
	NewValue("42").CompileTo(s.mapper)

	s.Assert().Empty(s.mapper.Vars)
	s.Assert().Equal(ast.NewValue("42"), s.mapper.Result)
}
