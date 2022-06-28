package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionString = ""
	Port             = 0
)

func Load() (int, string) {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	log.Println(os.Getenv("API_PORT"))
	APIPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		APIPort = 5000
	}
	Port = APIPort

	ConnectionString = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	return APIPort, ConnectionString
}
