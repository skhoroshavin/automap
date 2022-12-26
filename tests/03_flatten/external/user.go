package external

type User struct {
	ID        string
	FirstName string
	LastName  string
	Address   Address
}

type Address struct {
	Street  string
	City    string
	Country string
	Code    int
}
