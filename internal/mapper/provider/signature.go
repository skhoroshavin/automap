package provider

import (
	"github.com/skhoroshavin/automap/internal/mapper/types"
)

type Signature struct {
	Name Matcher
	Type *types.Type
}

func (s *Signature) Match(req *Request) bool {
	if req.Type.ID() != s.Type.ID() {
		return false
	}

	return s.Name.Match(req.Name)
}

func (s *Signature) Unpack(parent Provider) (res List) {
	for _, field := range s.Type.Fields {
		res = append(res, NewField(parent, &field))
	}
	for _, method := range s.Type.Methods {
		if method.ReturnType == nil {
			continue
		}
		res = append(res, NewMethod(parent, &method))
	}
	return res
}
