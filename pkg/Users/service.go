package Users

type Repository interface{
	FindAllUsers() ([]User, error)
}

type Service interface{
	GetUsers() ([]User, error)
}

type service struct{
	r Repository
}

//CREATE NEW SERVICE
func NewService(r Repository) Service {
	return &service{r}
}

//METHODS SERVICES
func(s *service) GetUsers() ([]User,error){
	return s.r.FindAllUsers()
}
