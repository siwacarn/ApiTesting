package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Indicator struct {
	gorm.Model
	Indicator   string
	Light       int     `json:"lightintensity"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func IndicatorDBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Indicator{})
	return db
}
