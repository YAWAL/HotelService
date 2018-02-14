package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

//var (
//	once sync.Once
//)

func DBconnection() (db *gorm.DB, err error) {
	//once.Do(func() {
	db, err = gorm.Open("postgres", "user=postgres dbname=hoteldb sslmode=disable password=root")
	if err != nil {
		log.Fatalf("Error during connection to Postgres has occurred %s", err.Error())
	} else {
		log.Println("connection to Pg has been established")
	}
	//})
	return db, err
}
