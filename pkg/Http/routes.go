package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/luis16121013/goApi/pkg/Repository"
	"github.com/luis16121013/goApi/pkg/Users"
)

func StartServer() {

	r,err := Repository.SetupRepositoryMysql()
	if err !=nil {
		fmt.Println(err)
	}

	server := echo.New()
	s := Users.NewService(r)

	server.GET("/api/v1", Users.AllUsers(s) )

	server.Logger.Fatal(server.Start( GetPort() ))
}

