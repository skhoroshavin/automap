package oldmapper

import (
	"github.com/skhoroshavin/automap/internal/parser"
)

func New(parseResults *parser.Result) (res *Registry, err error) {
	res = new(Registry)
	res.pkg = parseResults.Package

	res.imports = make([]string, 0, len(parseResults.Imports))
	for v := range parseResults.Imports {
		res.imports = append(res.imports, v)
	}

	res.mappers = make([]*Mapper, len(parseResults.Mappers))
	for i, v := range parseResults.Mappers {
		res.mappers[i], err = NewMapper(v, parseResults.TypeInfo)
		if err != nil {
			return
		}
	}

	return
}

type Registry struct {
	pkg     string
	imports []string
	mappers []*Mapper
}

func (r *Registry) Package() string {
	return r.pkg
}

func (r *Registry) Imports() []string {
	return r.imports
}

func (r *Registry) Mappers() []*Mapper {
	return r.mappers
}
