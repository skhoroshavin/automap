// Code generated by automap. DO NOT EDIT.

//go:build !automap

//go:generate automap

package my

import (
	"github.com/skhoroshavin/automap/tests/03_flatten/external"
)

func ValueToValue(user external.User) UserDTO {
	return UserDTO{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Country: user.Address.Country,
	}
}