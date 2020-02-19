package database

import (
	"log"
	"os"

	"github.com/blogster/models"
)

// DB used for global declaration og gorm.DB pointer
var DB *gorm.DB

// GetDB connects to db and passes out a global gorm.DB pointer
func GetDB() *gorm.DB {
	db, err = gorm.Open("postgres", "postgres:postgres@tcp(db:5432)/sample?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("error connecting to database : ", err)
		os.Exit(-1)
	}
	DB = db
	return DB
}

// Migrate defines the schema migrations
func Migrate() {
	db := GetDB()
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.React{})
}

// Apply defines foreign key relations
func Apply() {
	db := GetDb()
	db.Model(&models.Comment{}).AddForeignKey("post_id", "posts(id)", "CASCADE", "CASCADE")
	db.Model(&models.Comment{}).AddForeignKey("author_id", "authors(id)", "CASCADE", "CASCADE")
	db.Model(&models.Post{}).AddForeignKey("author_id", "authors(id)", "CASCADE", "CASCADE")
	db.Model(&models.React{}).AddForeignKey("author_id", "authors(id)", "CASCADE", "CASCADE")
	db.Model(&models.React{}).AddForeignKey("post_id", "posts(id)", "CASCADE", "CASCADE")
}
