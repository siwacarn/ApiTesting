package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"name"`
	Passwoed string `json:"password"`
	// TODO: need more data
}

// TODO: Revise function name
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
