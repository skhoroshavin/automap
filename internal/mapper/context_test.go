package mapper

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestContext(t *testing.T) {
	suite.Run(t, new(ContextSuite))
}

type ContextSuite struct {
	suite.Suite
	context *Context
}

func (s *ContextSuite) SetupSuite() {
	s.context = new(Context)
	s.context.AddProvider(&MockProvider{name: "x"})
	s.context.AddProvider(&MockProvider{name: "y"})
}

func (s *ContextSuite) TestResolveForEmptyContextFails() {
	_, err := new(Context).Resolve(Request{Name: "x"})
	s.Require().Error(err)
}

func (s *ContextSuite) TestResolveForExistingProviderSucceeds() {
	node, err := s.context.Resolve(Request{Name: "x"})
	s.Require().NoError(err)
	s.Require().NotNil(node)
}

func (s *ContextSuite) TestResolveForNonExistingProviderFails() {
	_, err := s.context.Resolve(Request{Name: "z"})
	s.Require().Error(err)
}
