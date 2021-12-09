package app

import (
	"github.com/nunonano/hacktiv-final-project/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=ec2-54-157-113-118.compute-1.amazonaws.com user=zvmehrtxtwnmea password=4af7ba68f834dba2c0fccfc4cc8f348b09f2b52050fb47b532afcb5bbed1a61e dbname=d1frg8fdqfata0 port=5432 TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Debug().AutoMigrate(entity.Order{}, entity.Todo{})

	return db
}
