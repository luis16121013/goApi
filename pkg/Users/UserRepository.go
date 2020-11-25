package Users

type UserRepository interface {
	FindAllUsers() ([]User, error)
}
