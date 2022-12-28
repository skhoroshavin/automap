package mapper

type Request struct {
	Name string
	Type *Type
}

func (r *Request) TypeCasts() (res ProviderList) {
	res = ProviderList{}

	if r.Type.IsPointer {
		// Add deref provider
	} else {
		// Add ref provider
	}

	if r.Type.IsStruct {
		// Add struct builder provider
	}

	return res
}
