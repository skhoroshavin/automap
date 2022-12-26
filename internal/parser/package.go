package parser

import (
	"golang.org/x/tools/go/packages"
	"strings"
)

type Package struct {
	Name string
	Path string
}

func ParsePackage(pkg *packages.Package) *Package {
	return &Package{
		Name: parsePackageName(pkg.ID),
		Path: pkg.ID,
	}
}

func parsePackageName(path string) string {
	idx := strings.LastIndex(path, "/")
	if idx < 0 {
		return path
	}
	return path[idx+1:]
}
