package parser

import (
	"github.com/skhoroshavin/automap/internal/mapper"
	"go/types"
)

func BuildPackageConfig(parseResults *Result) (res *mapper.PackageConfig, err error) {
	res = new(mapper.PackageConfig)
	res.Name = parseResults.Package

	res.Imports = make([]string, 0, len(parseResults.Imports))
	for v := range parseResults.Imports {
		res.Imports = append(res.Imports, v)
	}

	res.Mappers = make([]*mapper.Config, len(parseResults.Mappers))
	for i, v := range parseResults.Mappers {
		res.Mappers[i], err = BuildMapperConfig(v, parseResults.TypeInfo)
		if err != nil {
			return
		}
	}

	return
}

func BuildMapperConfig(src *Mapper, typeInfo *types.Info) (res *mapper.Config, err error) {
	res = new(mapper.Config)
	res.Name = src.Name
	res.FromName = src.From.Name
	res.FromType, err = ParseType(src.From.Type, typeInfo)
	if err != nil {
		return
	}
	res.ToType, err = ParseType(src.To.Type, typeInfo)
	if err != nil {
		return
	}
	return
}
