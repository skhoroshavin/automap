//go:build automap

package test

import (
	"github.com/skhoroshavin/automap"
	"github.com/skhoroshavin/automap/internal/parser/test/another"
	some "github.com/skhoroshavin/automap/internal/parser/test/whatever"
)

func ValueToPtr(user another.User) *UserName {
	panic(automap.Build())
}

func PtrToValue(user *another.User) some.User {
	panic(automap.Build())
}

func NotMapper(user another.User) UserName {
	panic("not a mapper")
}
