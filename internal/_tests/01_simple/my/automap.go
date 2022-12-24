//go:build automap

package my

import (
	"automap"
	"automap/internal/_tests/01_simple/external"
)

func MapUserName(user *external.User) *UserName {
	panic(automap.Build())
}
