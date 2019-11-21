package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"name"`
	Passwoed string `json:"password"`
}

type Indicator struct {
	gorm.Model
	Light       int     `json:"light"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func DBMigrate(db *gorm) *gorm.DB {
	db.AutoMigrate(&Indicator{})
	return db
}
