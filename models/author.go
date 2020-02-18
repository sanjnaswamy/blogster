package models

// Author defines schema for user on blog website
type Author struct {
	ID       int    `gorm:"type:int;not null;auto_increment;primary_key"`
	Name     string `gorm:"type:varchar(24);not null;`
	UserName string `gorm:"type:varchar(24);not null;unique"`
	Email    string `gorm:"type:varchar(24);not null;"`
	Password string `gorm:"type:varchar(255);not null"`
}
