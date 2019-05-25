package main

import (
	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"

	_ "github.com/pajbot/pajbot2/docs" // docs is generated by Swag CLI, you have to import it.
)

func main() {
	e := echo.New()

	e.GET("/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
