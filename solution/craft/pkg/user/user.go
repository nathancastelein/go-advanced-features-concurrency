package user

type User struct {
	Firstname string
	Lastname  string
}

type Lister interface {
	List() ([]User, error)
}
