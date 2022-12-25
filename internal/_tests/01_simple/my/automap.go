//go:build automap

package my

import (
	"automap"
	"automap/internal/_tests/01_simple/external"
)

func ValueToValue(user external.User) UserName {
	panic(automap.Build())
}

func ValueToPtr(user external.User) *UserName {
	panic(automap.Build())
}

func PtrToValue(user *external.User) UserName {
	panic(automap.Build())
}

func PtrToPtr(user *external.User) *UserName {
	panic(automap.Build())
}
