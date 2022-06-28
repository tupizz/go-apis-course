package main

import (
	"github.com/tupizz/go-gin-api/database"
	"github.com/tupizz/go-gin-api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
