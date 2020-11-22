package Users

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func AllUsers(s Service) ( func(c echo.Context) error ){
	return func(c echo.Context) error{
		users, err := s.GetUsers()
		if err != nil{
			return c.JSON(http.StatusOK,users)
		}
		return c.JSON(http.StatusNotFound,users)
	}
}
