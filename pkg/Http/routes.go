package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/luis16121013/goApi/pkg/Repository"
	"github.com/luis16121013/goApi/pkg/Users"
)

func StartServer() {

	r, err := Repository.SetupRepositoryMysql()
	if err != nil {
		fmt.Println(err)
	}

	server := echo.New()
	s := Users.NewService(r)

	gAdmin := server.Group("/api/v1/admin")
	gAdmin.GET("/hola", func(c echo.Context) error {
		return c.HTML(200, "<h1>HOLA DESDE EL GROUP!</h1>")
	})

	server.GET("/api/v1", Users.AllUsers(s))

	server.Logger.Fatal(server.Start(GetPort()))
}
