//go:build automap

package my

import (
	"github.com/skhoroshavin/automap"
	"github.com/skhoroshavin/automap/tests/03_flatten/external"
)

func ValueToValue(user external.User) UserDTO {
	panic(automap.Build())
}
