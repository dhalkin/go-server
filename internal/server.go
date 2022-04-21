package internal

import (
	"fmt"
	"github.com/dhalkin/go-server/internal/controllers"
	"github.com/dhalkin/go-server/internal/projectpath"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func StartServer() {

	// Search config in home directory with name ".go-game" (without extension).
	viper.AddConfigPath(projectpath.Root + "/configs")
	viper.SetConfigType("yaml")
	viper.SetConfigName("joker")

	viper.AutomaticEnv() // read in environment variables that match

	//there are no keys, why?
	//fmt.Println("viper allKeys:", viper.AllKeys())
	//fmt.Println("project path:", projectpath.Root)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	e := getE()
	address := fmt.Sprintf("%v", viper.Get("Server.Addr"))

	if err := e.Start(address); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func getE() *echo.Echo {

	e := echo.New()
	e.HideBanner = false
	e.Debug = true
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.BodyLimit("1M"))
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {

		// do not log body if multipart data
		req := c.Request()
		err := req.ParseMultipartForm(1024)
		if err != nil {
			str := string(reqBody)
			fmt.Println(str)
		}
		fmt.Println(string(resBody))
	}))

	// define db

	// add middleware and routes
	e.GET("/hello", controllers.Hello)
	e.GET("/hola", controllers.GetHola)
	e.POST("/hola", controllers.PutHola)
	e.GET("/container", controllers.GetContainer)

	return e
}
