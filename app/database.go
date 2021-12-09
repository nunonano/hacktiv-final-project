package app

import (
	"github.com/nunonano/hacktiv-final-project/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=db user=postgres password=secret dbname=hacktiv_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Debug().AutoMigrate(entity.Order{}, entity.Todo{})

	return db
}
