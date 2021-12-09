package app

import (
	"fmt"

	"github.com/nunonano/hacktiv-final-project/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "ec2-54-157-113-118.compute-1.amazonaws.com"
	dbName   = "d1frg8fdqfata0"
	user     = "zvmehrtxtwnmea"
	password = "4af7ba68f834dba2c0fccfc4cc8f348b09f2b52050fb47b532afcb5bbed1a61e"
	port     = "5432"
)

func InitDB() *gorm.DB {
	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// port := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Debug().AutoMigrate(entity.Order{}, entity.Todo{})

	return db
}
