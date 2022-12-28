package mapper

import (
	"github.com/skhoroshavin/automap/internal/mapper/provider"
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
	s.context.AddProvider(provider.NewMock("x"))
	s.context.AddProvider(provider.NewMock("y", provider.NewMock("z")))
}

func (s *ContextSuite) TestResolveForEmptyContextFails() {
	_, err := new(Context).Resolve(provider.NewMockRequest("x"))
	s.Require().Error(err)
}

func (s *ContextSuite) TestResolveForExistingProviderSucceeds() {
	node, err := s.context.Resolve(provider.NewMockRequest("x"))
	s.Require().NoError(err)
	s.Require().NotNil(node)
}

func (s *ContextSuite) TestResolveForNonExistingProviderFails() {
	_, err := s.context.Resolve(provider.NewMockRequest("a"))
	s.Require().Error(err)
}

func (s *ContextSuite) TestResolveForNestedProviderSucceeds() {
	node, err := s.context.Resolve(provider.NewMockRequest("z"))
	s.Require().NoError(err)
	s.Require().NotNil(node)
}
