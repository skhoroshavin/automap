package provider

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestRequest(t *testing.T) {
	suite.Run(t, new(RequestSuite))
}

type RequestSuite struct {
	suite.Suite
	providers List
}

func (s *RequestSuite) SetupSuite() {
	s.providers = List{
		NewMock("x"),
		NewMock("y",
			NewMock("z"),
		),
	}
}

func (s *RequestSuite) TestResolveForEmptyContextFails() {
	_, err := NewMockRequest("x").Resolve(nil)
	s.Require().Error(err)
}

func (s *RequestSuite) TestResolveForExistingProviderSucceeds() {
	node, err := NewMockRequest("x").Resolve(s.providers)
	s.Require().NoError(err)
	s.Require().NotNil(node)
}

func (s *RequestSuite) TestResolveForNonExistingProviderFails() {
	_, err := NewMockRequest("a").Resolve(s.providers)
	s.Require().Error(err)
}

func (s *RequestSuite) TestResolveForNestedProviderSucceeds() {
	node, err := NewMockRequest("z").Resolve(s.providers)
	s.Require().NoError(err)
	s.Require().NotNil(node)
}
