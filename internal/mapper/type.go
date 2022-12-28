package mapper

// Type represents generic type
type Type interface {
	// ID is a unique type identifier, which can be used for comparisons
	ID() string
	// Casts returns list of possible type casting providers
	Casts(string) ProviderList
}

type MockType struct {
	id string
}

func (m *MockType) ID() string {
	return m.id
}

func (m *MockType) Casts(name string) ProviderList {
	return ProviderList{}
}
