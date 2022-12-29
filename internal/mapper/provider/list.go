package provider

// List represents list of providers
type List []Provider

// ForEach does a BFS traversal of provider list until either all nodes are
// visited or visitor function returns true
func (l List) ForEach(v Visitor) {
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

type Visitor func(p Provider) bool
