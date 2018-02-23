package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"sync"
)

var (
	once sync.Once
)

func DBconnection() (db *gorm.DB, err error) {
	once.Do(func() {
	db, err = gorm.Open("postgres", "user=postgres dbname=hoteldb sslmode=disable password=root")
	if err != nil {
		log.Fatalf("Error during connection to PostgreSQL has occurred: %s", err.Error())
	} else {
		log.Println("Connection to PostgreSQL has been established")
	}
	})
	return db, err
}
