package main

import (
	"fmt"
	"reflect"

	"github.com/ank1106/webserver/src/main/models"
	"github.com/ank1106/webserver/src/main/server"
)

func main() {
	config := models.DBConfig{
		Type: "sqlite3",
		URL:  "test.db",
	}
	db := models.ConnectDB(&config)
	webserver := server.WebServer{}
	webserver.Run(&db)
	fmt.Println(webserver, reflect.TypeOf(webserver))

}
