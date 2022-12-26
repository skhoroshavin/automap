package external

type User struct {
	id        string
	firstName string
	lastName  string
}

func (u *User) ID() string {
	return u.id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}
