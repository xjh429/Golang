package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDB() (*gorm.DB, error) {
	dsn := "root:1234@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func GetDB() *gorm.DB {
	db, err := connectDB()
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
