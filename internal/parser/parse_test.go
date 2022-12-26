package parser

import (
	"fmt"
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
		`some "github.com/skhoroshavin/automap/internal/parser/test/whatever"`,
	}, s.cfg.Imports)
}

func (s *ParseSuite) TestMapperList() {
	mappers := lo.Map(s.cfg.Mappers, func(m *mapper.Config, _ int) string { return m.Name })
	s.Assert().Equal([]string{
		"ValueToPtr",
		"PtrToValue",
	}, mappers)
}

func (s *ParseSuite) TestMapperSignature() {
	m, ok := lo.Find(s.cfg.Mappers, func(m *mapper.Config) bool { return m.Name == "ValueToPtr" })
	s.Require().True(ok)

	s.Assert().Equal("user", m.FromName)
	s.Assert().Equal("another.User", m.FromType.Name())
	s.Assert().False(m.FromType.IsPointer())
	s.Assert().Equal("UserName", m.ToType.Name())
	s.Assert().True(m.ToType.IsPointer())
}

func (s *ParseSuite) TestUserType() {
	m, ok := lo.Find(s.cfg.Mappers, func(m *mapper.Config) bool { return m.Name == "ValueToPtr" })
	s.Require().True(ok)

	s.Require().Equal("another.User", m.FromType.Name())
	user, ok := m.FromType.(*mapper.StructType)
	s.Require().True(ok)

	userFields := lo.Map(user.Fields, func(p mapper.Provider, _ int) string {
		return fmt.Sprintf("%s %s", p.Name, p.Type.Name())
	})
	s.Assert().Equal([]string{
		"ID string",
		"FirstName string",
		"LastName string",
		"Address another.Address",
	}, userFields)
}
