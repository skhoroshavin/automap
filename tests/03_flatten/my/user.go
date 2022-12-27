package my

type UserDTO struct {
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	AddressCode   int    `json:"address_code,omitempty"`
	StreetAndCity string `json:"street_and_city,omitempty"`
	Country       string `json:"country,omitempty"`
	PrefsGolang   bool   `json:"prefs_golang,omitempty"`
	PrefsCompiled bool   `json:"prefs_compiled,omitempty"`
}
