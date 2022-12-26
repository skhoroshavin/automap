package oldmapper

import (
	"github.com/skhoroshavin/automap/internal/mapper"
	"github.com/skhoroshavin/automap/internal/parser"
	"go/types"
)

func NewMapper(mapper *parser.Mapper, typeInfo *types.Info) (res *Mapper, err error) {
	res = new(Mapper)
	res.Name = mapper.Name
	res.FromName = mapper.From.Name
	res.FromType, err = ParseType(mapper.From.Type, typeInfo)
	if err != nil {
		return
	}
	res.ToType, err = ParseType(mapper.To.Type, typeInfo)
	if err != nil {
		return
	}
	return
}

type Mapper struct {
	Name     string
	FromName string
	FromType mapper.Type
	ToType   mapper.Type
}
