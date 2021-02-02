package user

type UserService interface {
	CreateUser(user *User) (*Response, error)
	UpdateUser(user *User, id int) error
	DeleteUser(id int) error
	GetUserByEmail(email string) (user User, err error)
	GetUserById(id int) (user User, err error)
}


func NewUserService() (service UserService) {
	r := NewUserRepository()
	service = &Service{
		Repository: r,
	}
	return
}