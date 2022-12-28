package mapper

// ProviderList represents list of providers
type ProviderList []Provider

// ForEach does a BFS traversal of provider list until either all visitedNodes are visited
// or visitor function returns true
func (l ProviderList) ForEach(v ProviderVisitor) {
	q := make([]Provider, len(l))
	for i, p := range l {
		q[i] = p
	}

	for len(q) > 0 {
		p := q[0]
		if v(p) {
			return
		}

		q = q[1:]
		for _, cp := range p.Children() {
			q = append(q, cp)
		}
	}
}

type ProviderVisitor func(p Provider) bool
