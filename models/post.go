package models

// Post defines schema for posts on blog website
type Post struct {
	ID       int    `gorm:"type:int;not null;auto_increment;primary_key"`
	Content  string `gorm:"type:varchar;not null;"`
	Time     string `gorm:"type:time; not null"`
	AuthorID int    `gorm:"type:int;not null;foreignkey:AuthorID"`
}
