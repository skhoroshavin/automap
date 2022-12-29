package provider

import "strings"

type Matcher []string

func NewMatcher(names ...string) (res Matcher) {
	res = make(Matcher, 0, len(names))
	for _, name := range names {
		if len(name) == 0 {
			continue
		}
		res = append(res, name)
	}
	return
}

func (m Matcher) Match(name string) bool {
	if len(m) == 0 || name == "" {
		return true
	}

	for _, item := range m {
		if strings.ToLower(item) == strings.ToLower(name) {
			return true
		}
	}
	return false
}

func (m Matcher) Append(name string, allowGlobal bool) (res Matcher) {
	res = make(Matcher, len(m), len(m)+1)
	for i, base := range m {
		res[i] = base + name
	}
	if allowGlobal {
		res = append(res, name)
	}
	return res
}
