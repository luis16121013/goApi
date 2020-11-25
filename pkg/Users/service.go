package Users

type Service interface {
	GetUsers() ([]User, error)
}

type service struct {
	r UserRepository
}

//CREATE NEW SERVICE
func NewService(r UserRepository) Service {
	return &service{r}
}

//METHODS SERVICES
func (s *service) GetUsers() ([]User, error) {
	return s.r.FindAllUsers()
}
