package database

import (
	"assignment-2/entity"
	// "database/sql"
	"fmt"
	"log"

	// "time"

	// "github.com/google/uuid"
	// _ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "12345678"
	dbname   = "postgres"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(entity.Items{}, entity.Orders{})
}

func GetDB() *gorm.DB {
	return db
}
