package models

import (
	"github.com/jinzhu/gorm"
)

func init() {

}

func InitGorm(db *gorm.DB) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}

}
