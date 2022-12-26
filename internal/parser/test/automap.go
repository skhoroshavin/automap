//go:build automap

package test

import (
	"github.com/skhoroshavin/automap"
	"github.com/skhoroshavin/automap/internal/parser/test/another"
)

func ValueToValue(user another.User) UserName {
	panic(automap.Build())
}

func ValueToPtr(user another.User) *UserName {
	panic(automap.Build())
}

func PtrToValue(user *another.User) UserName {
	panic(automap.Build())
}

func PtrToPtr(user *another.User) *UserName {
	panic(automap.Build())
}
