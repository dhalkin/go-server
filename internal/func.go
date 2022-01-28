package internal

import (
	"fmt"
	"github.com/dhalkin/go-server/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getE() *echo.Echo {

	e := echo.New()
	e.HideBanner = false
	e.Debug = true
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.BodyLimit("1M"))
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {

		// ??
		//_, err := c.FormFile("file")
		//if err != nil {
		//	str := string(reqBody)
		//	fmt.Println(str)
		//}

		req := c.Request()
		mpfd := req.MultipartForm.File
		if len(mpfd) == 0 {
			str := string(reqBody)
			fmt.Println(str)
		}

		fmt.Println(string(resBody))

	}))

	// add middleware and routes
	e.GET("/hello", handlers.Hello)
	e.GET("/hola", handlers.GetHola)
	e.POST("/hola", handlers.PutHola)

	return e
}
