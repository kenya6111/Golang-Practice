package user

type Repository interface {
	Save(u *User) error
	FindAll() ([]*User, error)
}
