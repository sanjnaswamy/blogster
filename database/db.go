package database

import (
	"log"
	"os"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	db, err = gorm.Open("postgres", "postgres:postgres@tcp(db:5432)/sample?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("error connecting to database : ", err)
		os.Exit(-1)
	}
	DB = db
	return DB
}
