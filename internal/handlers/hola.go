package handlers

import (
	"fmt"
	"github.com/dhalkin/go-server/internal/projectpath"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
)

func GetHola(c echo.Context) error {

	return c.String(http.StatusOK, "OK - hola")
}

func PutHola(c echo.Context) error {

	path := fmt.Sprintf("%v", viper.Get("Sandbox.LocalFolder"))
	basePath := projectpath.Root + "/" + path

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(basePath + "/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK - Put")
}
