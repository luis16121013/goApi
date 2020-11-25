package Repository

import (
	"errors"
	"github.com/luis16121013/goApi/pkg/Users"
)

type Usrs []Users.User

func (s *RepositoryMysql) FindAllUsers() ([]Users.User, error) {
	users := Usrs{}
	sqlDB := "SELECT id,username,password,email FROM users"
	rows, err := s.Query(sqlDB)
	if err != nil {
		return users, errors.New("Users not found!")
	}

	for rows.Next() {
		u := Users.User{}
		rows.Scan(&u.Id, &u.Username, &u.Password, &u.Email)
		users = append(users, u)
	}
	return users, nil
}
