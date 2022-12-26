package external

type User struct {
	ID        string
	FirstName string
	LastName  string
	Address   Address
	prefs     Prefs
}

func (u *User) Prefs() *Prefs {
	return &u.prefs
}

type Address struct {
	Street  string
	City    string
	Country string
	Code    int
}

func (a *Address) StreetAndCity() string {
	return ""
}

type Prefs struct {
	Golang     bool
	Typescript bool
}

func (p *Prefs) Compiled() bool {
	return false
}
