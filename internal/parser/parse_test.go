package parser

import (
	"github.com/samber/lo"
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

func (s *ParseSuite) TestMapperList() {
	mappers := lo.Map(s.cfg.Mappers, func(m *mapper.Config, _ int) string { return m.Name })
	s.Assert().ElementsMatch([]string{
		"PtrToValue",
		"ValueToPtr",
	}, mappers)
}

func (s *ParseSuite) TestMapperSignature() {
	m, _ := lo.Find(s.cfg.Mappers, func(m *mapper.Config) bool { return m.Name == "ValueToPtr" })
	s.Assert().Equal("user", m.FromName)
	s.Assert().Equal("another.User", m.FromType.Name())
	s.Assert().False(m.FromType.IsPointer())
	s.Assert().Equal("UserName", m.ToType.Name())
	s.Assert().True(m.ToType.IsPointer())
}
