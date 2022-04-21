package main

import (
	"github.com/dhalkin/go-server/internal"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	internal.StartServer()
}
