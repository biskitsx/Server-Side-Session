package database

import (
	"github.com/biskitsx/Server-Side-Session/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectPostgres() error {
	var err error
	dsn := "host=localhost user=postgres password=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	Db.AutoMigrate(&model.User{})
	return nil
}
