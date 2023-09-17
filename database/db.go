package database

import (
	"log"

	"github.com/biskitsx/Server-Side-Session/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return db
	}
	db.AutoMigrate(&model.User{})
	return db
}
