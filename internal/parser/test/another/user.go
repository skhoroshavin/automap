package another

type User struct {
	ID        string
	FirstName string
	LastName  string
	Address   Address
}

type Address struct {
	Street   string
	City     string
	PostCode int
	Country  string
}
