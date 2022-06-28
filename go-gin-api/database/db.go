package database

import (
	"fmt"
	"log"

	"github.com/tupizz/go-gin-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

type DbConnection struct {
	user   string
	pass   string
	host   string
	port   int
	dbName string
}

var db = DbConnection{
	user:   "root",
	pass:   "root",
	host:   "localhost",
	port:   5432,
	dbName: "root",
}

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		db.host,
		db.user,
		db.pass,
		db.dbName,
		db.port,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("error when connecting to database")
	}

	DB.AutoMigrate(&models.Student{})
}
