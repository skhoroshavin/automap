package parser

import (
	"github.com/skhoroshavin/automap/internal/mapper"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestParse(t *testing.T) {
	suite.Run(t, new(ParseSuite))
}

type ParseSuite struct {
	suite.Suite
	cfg *mapper.PackageConfig
}

func (s *ParseSuite) SetupSuite() {
	cfg, err := Parse("test")
	s.Require().NoError(err)
	s.cfg = cfg
}

func (s *ParseSuite) TestPackageName() {
	s.Assert().Equal("test", s.cfg.Name)
}

func (s *ParseSuite) TestImports() {
	s.Assert().Equal([]string{
		`"github.com/skhoroshavin/automap/internal/parser/test/another"`,
	}, s.cfg.Imports)
}
