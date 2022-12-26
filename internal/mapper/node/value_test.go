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
}

func (s *ValueSuite) TestValueBuildsSimpleReturnStatement() {
	body := build(NewValue("42"))

	s.Assert().Empty(body.Vars)
	s.Assert().Equal(ast.NewValue("42"), body.Result)
}
