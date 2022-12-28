package mapper

type Request struct {
	Name string
	Type Type
}

func (r *Request) TypeCasts() ProviderList {
	return r.Type.Casts(r.Name)
}
