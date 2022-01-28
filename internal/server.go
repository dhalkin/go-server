package internal

import (
	"fmt"
	"github.com/dhalkin/go-server/internal/projectpath"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func StartServer() {

	configure()
	e := getE()
	address2 := fmt.Sprintf("%v", viper.Get("Server.Addr"))

	if err := e.Start(address2); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func configure() {

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
}
