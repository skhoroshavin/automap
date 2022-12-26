package my

type UserDTO struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	//AddressCode int    `json:"address_code,omitempty"`
	Country string `json:"country,omitempty"`
}
