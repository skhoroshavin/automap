package provider

import (
	"github.com/skhoroshavin/automap/internal/mapper/types"
)

type Request struct {
	Name string
	Type *types.Type
}

func (r *Request) TypeCasts() (res List) {
	if r.Type.IsPointer {
		// Add deref provider
	} else {
		// Add ref provider
	}

	if r.Type.IsStruct {
		// Add struct builder provider
	}

	return
}
