package main

import (
	"fmt"

	"github.com/tupizz/go-api-rest/database"
	"github.com/tupizz/go-api-rest/routes"
)

func main() {
	database.Connect()
	fmt.Println("initing server with go")
	routes.HandleRequest()
}
