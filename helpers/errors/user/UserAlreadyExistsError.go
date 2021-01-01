package user

type AlreadyExists struct {}


func (AlreadyExists) Error() string{
	return "An user with this email already exists"
}