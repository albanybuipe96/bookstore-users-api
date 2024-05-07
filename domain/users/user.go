package users

type User struct {
	Id       string
	Username string
	Email    string
	Password string
}

func (user *User) Validate() error {
	return nil
}
