package users

type User struct {
	Id          int64  `db:"id"`
	FirstName   string `db:"firstname"`
	LastName    string `db:"lastname"`
	Email       string `db:"email"`
	DateCreated string `db:"datecreated"`
}

func (user *User) Validate() error {
	return nil
}
