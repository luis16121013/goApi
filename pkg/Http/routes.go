package router

import (
	"os"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/luis16121013/apiGo/pkg/Repository"
)

func StartServer() {
	
	server := echo.New()

	r,err := Repository.SetupRepositoryMysql()
	if err !=nil {
		fmt.Println(err)
	}

	server.GET("/", func(c echo.Context) error{
		return c.String(200, "ok!")
	})

	server.Logger.Fatal(server.Start( getPort() ))
}


func getPort() string{
	port := os.Getenv("PORT")
	if port != ""{
		return ":"+port
	}
	return ":3000"
}
