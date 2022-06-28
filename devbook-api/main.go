package main

import (
	"devbook-api/src/config"
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Println(config.ConnectionString)
	fmt.Println(config.Port)

	r := router.Gerar()
	log.Printf("Running application on port http://localhost:%d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
