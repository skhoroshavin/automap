package provider

import (
	"github.com/skhoroshavin/automap/internal/mapper/types"
	"strings"
)

type Signature struct {
	Name string
	Type *types.Type
}

func (s *Signature) Match(req *Request) bool {
	if req.Type.ID() != s.Type.ID() {
		return false
	}

	if (req.Name == "") || (s.Name == "") {
		return true
	}

	return strings.ToLower(req.Name) == strings.ToLower(s.Name)
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
