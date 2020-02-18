package models

// React defines schema for reactions on blog posts
type React struct {
	ID       int    `gorm:"type:int;not null;auto_increment;primary_key"`
	Action   string `gorm:"type:varchar(10);not null"`
	AuthorID int    `gorm:"type:int;not null;foreignkey:AuthorID"`
	PostID   int    `gorm:"type:int;not null;foreignkey:PostID"`
}
