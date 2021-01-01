package user

type CannotCreateError struct{}

func (CannotCreateError) Error() string {
	return "Something gone wrong in user creation process"
}
